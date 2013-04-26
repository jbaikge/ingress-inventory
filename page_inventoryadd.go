package main

import (
	"github.com/jbaikge/ingress-inventory/parser"
	"log"
	"net/http"
)

func init() {
	router.HandleFunc("/inventory/update", AuthWrapper(HandleInventoryUpdate))
	router.HandleFunc("/inventory/save", AuthWrapper(HandleInventorySave))
}

func HandleInventoryUpdate(w http.ResponseWriter, r *http.Request, ctx *parser.Context) {
	ctx.Title = "Ingress Inventory"
	if err := parser.Render(w, ctx, "update-inventory.html"); err != nil {
		log.Println(err)
	}
}

func HandleInventorySave(w http.ResponseWriter, r *http.Request, ctx *parser.Context) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/inventory", http.StatusTemporaryRedirect)
		return
	}
	http.Redirect(w, r, "/inventory", http.StatusTemporaryRedirect)
}
