package main

import (
	"github.com/jbaikge/ingress-inventory/parser"
	"log"
	"net/http"
)

func init() {
	router.HandleFunc("/inventory-add", AuthWrapper(HandleInventoryAdd))
}

func HandleInventoryAdd(w http.ResponseWriter, r *http.Request, ctx *parser.Context) {
	ctx.Title = "Ingress Inventory"
	if err := parser.Render(w, ctx, "update-inventory.html"); err != nil {
		log.Println(err)
	}
}
