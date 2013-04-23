package main

import (
	"github.com/jbaikge/ingress-inventory/parser"
	"github.com/jbaikge/ingress-inventory/profile"
	"log"
	"net/http"
)

type AuthFunc func(http.ResponseWriter, *http.Request, *parser.Context)

func AuthWrapper(f AuthFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		p := profile.Profile{}
		if cookie, err := r.Cookie("Profile"); err == nil {
			
			if err = sCookie.Decode(cookie.Name, cookie.Value, &p); err != nil {
				log.Print(err)
			}
		}

		log.Printf("Username: %s", p.DisplayUsername)

		ctx := &parser.Context{
			Profile: p,
		}
		f(w, r, ctx)
	}
}
