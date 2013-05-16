package main

import (
	"github.com/gorilla/schema"
	"github.com/jbaikge/ingress-inventory/inventory"
	"github.com/jbaikge/ingress-inventory/parser"
	"log"
	"net/http"
	"reflect"
	"strconv"
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
	inventoryDecoder.RegisterConverter(inventory.Rarity(0), convertRarity)
	inventoryDecoder.RegisterConverter(inventory.Mods{}, convertRareItems)
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
	log.Printf("%+v", s)
	http.Redirect(w, r, "/inventory", http.StatusTemporaryRedirect)
}

func convertDatetime(value string) reflect.Value {
	if v, err := time.Parse(timeLayout, value); err == nil {
		return reflect.ValueOf(v)
	}
	log.Printf("Could not parse as time: '%s' using '%s'", value, timeLayout)
	return reflect.Value{}
}

func convertRarity(value string) reflect.Value {
	if v, err := strconv.ParseInt(value, 10, 0); err == nil {
		return reflect.ValueOf(inventory.Rarity(v))
	}
	return reflect.Value{}
}

func convertRareItems(value string) reflect.Value {
	log.Print(value)
	if v, err := strconv.ParseInt(value, 10, 0); err == nil {
		return reflect.ValueOf(inventory.Mods{0: int(v)})
	}
	return reflect.Value{}
}
