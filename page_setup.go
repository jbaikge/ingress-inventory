package main

import (
	"github.com/gorilla/schema"
	"github.com/jbaikge/ingress-inventory/communities"
	db "github.com/jbaikge/ingress-inventory/mongo"
	"github.com/jbaikge/ingress-inventory/parser"
	"github.com/jbaikge/ingress-inventory/profile"
	"log"
	"net/http"
)

type setupExtra struct {
	Errors map[string]string
}

var setupDecoder *schema.Decoder

func init() {
	router.HandleFunc("/setup", HandleSetup)
	router.HandleFunc("/setupThanks", HandleSetupThanks)
	setupDecoder = schema.NewDecoder()
}

func HandleSetup(w http.ResponseWriter, r *http.Request) {
	// Grab Profile from cookie set in HandleLoginOAuth
	p := profile.Profile{}
	if cookie, err := r.Cookie("Profile"); err == nil {
		if err = sCookie.Decode(cookie.Name, cookie.Value, &p); err != nil {
			log.Print(err)
		}
	}
	// Make sure we got the profile out of the cookie
	if p.GoogleId == "" {
		http.Redirect(w, r, "/cannotGetCookie", http.StatusTemporaryRedirect)
		return
	}

	p.Communities = communities.All()

	e := newExtra()

	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			log.Println(err)
		}
		// Fill the profile communities with all, filter out falses later
		if err := setupDecoder.Decode(&p, r.PostForm); err != nil {
			log.Println(err)
		}
		if err := validateUsername(&p); err != "" {
			e.Errors["DisplayUsername"] = err
		}
		if err := validateCommunities(&p); err != "" {
			e.Errors["Communities"] = err
		}

		if len(e.Errors) == 0 {
			http.Redirect(w, r, "/setupThanks", http.StatusTemporaryRedirect)
			return
		}
	}

	ctx := &parser.Context{
		Title:       "Setup Your Account",
		Description: "Establish your account details",
		Profile:     p,
		Extra:       e,
	}

	if err := parser.Render(w, ctx, "setup.html"); err != nil {
		log.Println(err)
	}
}

func HandleSetupThanks(w http.ResponseWriter, r *http.Request) {
	// Grab Profile from cookie set in HandleLoginOAuth
	p := profile.Profile{}
	if cookie, err := r.Cookie("Profile"); err == nil {
		if err = sCookie.Decode(cookie.Name, cookie.Value, &p); err != nil {
			log.Print(err)
		}
	}

	ctx := &parser.Context{
		Title:       "Thank You for Setting Up Your Account",
		Description: "",
		Profile:     p,
	}

	if err := parser.Render(w, ctx, "setup-thanks.html"); err != nil {
		log.Println(err)
	}
}

func newExtra() (e setupExtra) {
	return setupExtra{
		Errors: map[string]string{},
	}
}

func validateUsername(p *profile.Profile) (e string) {
	if p.DisplayUsername == "" {
		return "Please provide your username"
	}

	reg, err := db.UsernameRegistered(p.DisplayUsername)
	if err != nil {
		log.Print(err)
		return "Could not verify username"
	}
	if reg {
		return "Username already registered"
	}
	return
}

func validateCommunities(p *profile.Profile) (e string) {
	var onePicked bool
	for _, c := range p.Communities {
		if c.Selected {
			onePicked = true
		}
	}
	if !onePicked {
		return "Please select at least one community"
	}
	return
}
