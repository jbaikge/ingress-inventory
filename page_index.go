package main

import (
	"github.com/jbaikge/ingress-inventory/template"
	"net/http"
)

func init() {
	router.HandleFunc("/", HandleIndex)
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	ctx := &template.Context{
		Title:       "Ingress Inventory",
		Description: "Track your inventory",
	}
	template.WriteHeader(w, ctx)
	template.WriteFooter(w, ctx)
}
