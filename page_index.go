package main

import (
	"github.com/jbaikge/ingress-inventory/parser"
	"log"
	"net/http"
)

func init() {
	router.HandleFunc("/", HandleIndex)
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	ctx := &parser.Context{
		Title:       "Ingress Inventory",
		Description: "Track your inventory",
	}
	if err := parser.Render(w, ctx, "index.html"); err != nil {
		log.Println(err)
	}
}
