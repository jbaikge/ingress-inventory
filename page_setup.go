package main

import (
	"github.com/jbaikge/ingress-inventory/parser"

	"log"
	"net/http"
)

func init() {
	router.HandleFunc("/setup", HandleSetup)
}

func HandleSetup(w http.ResponseWriter, r *http.Request) {
	ctx := &parser.Context{
		Title:       "Setup Your Account",
		Description: "Connect your account with Google and ",
	}
	if err := parser.Render(w, ctx, "setup.html"); err != nil {
		log.Println(err)
	}
}
