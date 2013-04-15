package main

import (
	"github.com/jbaikge/ingress-inventory/parser"
	"log"
	"net/http"
)

func init() {
	router.HandleFunc("/profileNotFound", HandleTrouble)
	router.HandleFunc("/cannotSetCookie", HandleTrouble)
}

func HandleTrouble(w http.ResponseWriter, r *http.Request) {
	ctx := &parser.Context{
		Title:       "Uh-oh",
		Description: "Something went wrong",
	}
	if err := parser.Render(w, ctx, "trouble.html"); err != nil {
		log.Println(err)
	}
}
