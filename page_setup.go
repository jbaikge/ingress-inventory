package main

import (
	"github.com/jbaikge/ingress-inventory/parser"
	"github.com/jbaikge/ingress-inventory/profile"

	"log"
	"net/http"
)

func init() {
	router.HandleFunc("/setup", HandleSetup)
}

func HandleSetup(w http.ResponseWriter, r *http.Request) {
	// Grab Profile from cookie set in HandleLoginOAuth
	p := profile.Profile{}
	if cookie, err := r.Cookie("Profile"); err == nil {
		if err = sCookie.Decode(cookie.Name, cookie.Value, &p); err != nil {
			log.Print(err)
		}
	}
	if p.Id == "" {
		http.Redirect(w, r, "/cannotGetCookie", http.StatusTemporaryRedirect)
		return
	}

	ctx := &parser.Context{
		Title:       "Setup Your Account",
		Description: "Establish your account details",
		Profile:     p,
	}
	if err := parser.Render(w, ctx, "setup.html"); err != nil {
		log.Println(err)
	}
}
