package main

import (
	"github.com/jbaikge/ingress-inventory/parser"
	"log"
	"net/http"
)

type AuthFunc func(http.ResponseWriter, *http.Request, *parser.Context)

func AuthWrapper(f AuthFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := &parser.Context{}

		cookie, err := r.Cookie("Id")
		if err != nil {
			log.Printf("AuthWrapper: %s", err) // "Named cookie not present"
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		if err = sCookie.Decode(cookie.Name, cookie.Value, &ctx.Profile.Id); err != nil {
			log.Print(err)
			// TODO: redirect to a 500 page or clear cookies and bounce to /
		}
		log.Printf("Cookie ID: %s", ctx.Profile.Id)
		// TODO: use id to fetch ctx.Profile

		f(w, r, ctx)
	}
}
