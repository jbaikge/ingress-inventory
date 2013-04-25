package main

import (
	"github.com/jbaikge/ingress-inventory/parser"
	"labix.org/v2/mgo/bson"
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

		var id bson.ObjectId
		if err = sCookie.Decode(cookie.Name, cookie.Value, &id); err != nil {
			log.Print(err)
			// TODO: redirect to a 500 page or clear cookies and bounce to /
		}
		// TODO: use id to fetch ctx.Profile

		f(w, r, ctx)
	}
}
