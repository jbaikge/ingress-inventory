package main

import (
	"github.com/gorilla/schema"
	"github.com/jbaikge/ingress-inventory/inventory"
	"github.com/jbaikge/ingress-inventory/mongo"
	"github.com/jbaikge/ingress-inventory/parser"
	"log"
	"net/http"
	"reflect"
	"time"
)

var (
	inventoryDecoder *schema.Decoder
	timeLayout       = time.RFC3339
)

func init() {
	router.HandleFunc("/inventory/update", AuthWrapper(HandleInventoryUpdate))
	router.HandleFunc("/inventory/save", AuthWrapper(HandleInventorySave))
	inventoryDecoder = schema.NewDecoder()
	inventoryDecoder.RegisterConverter(time.Time{}, convertDatetime)
}

func HandleInventoryUpdate(w http.ResponseWriter, r *http.Request, ctx *parser.Context) {
	ctx.Title = "Ingress Inventory"
	ctx.AddJavascript("/js/update-inventory.js")
	if err := parser.Render(w, ctx, "update-inventory.html"); err != nil {
		log.Println(err)
	}
}

func HandleInventorySave(w http.ResponseWriter, r *http.Request, ctx *parser.Context) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/inventory", http.StatusTemporaryRedirect)
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	s := inventory.NewState()
	// Fill the profile communities with all, filter out falses later
	if err := inventoryDecoder.Decode(&s, r.PostForm); err != nil {
		errs, _ := err.(schema.MultiError)
		for f, e := range errs {
			log.Printf("%14s: %s", f, e)
		}
	}
	log.Printf("%+v\n%+v", ctx.Profile, s)
	if err := mongo.AddInventory(ctx.Profile.Id, &s); err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/inventory/update", http.StatusTemporaryRedirect)
}

func convertDatetime(value string) reflect.Value {
	if v, err := time.Parse(timeLayout, value); err == nil {
		return reflect.ValueOf(v)
	}
	log.Printf("Could not parse as time: '%s' using '%s'", value, timeLayout)
	return reflect.Value{}
}
