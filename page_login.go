package main

import (
	"code.google.com/p/goauth2/oauth"
	"code.google.com/p/google-api-go-client/plus/v1"
	"github.com/jbaikge/ingress-inventory/profile"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	router.HandleFunc("/login", HandleLogin)
	router.HandleFunc("/loginOAuth", HandleLoginOAuth)
}

var config = &oauth.Config{
	ClientId:     "180854220287-c47islde6hggldt91sq5aeta7m3eenhf.apps.googleusercontent.com",
	ClientSecret: os.Getenv("SECRET"),
	Scope:        plus.PlusMeScope,
	AuthURL:      "https://accounts.google.com/o/oauth2/auth",
	TokenURL:     "https://accounts.google.com/o/oauth2/token",
	RedirectURL:  "http://localhost:8080/loginOAuth",
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	code := time.Now().Format(time.RFC3339Nano)
	encoded, err := sCookie.Encode("Code", code)
	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/cannotSetCookie", http.StatusTemporaryRedirect)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "Code",
		Value: encoded,
		Path:  "/",
	})
	http.Redirect(w, r, config.AuthCodeURL(code), http.StatusFound)
}

func HandleLoginOAuth(w http.ResponseWriter, r *http.Request) {
	var code string
	var err error

	// Grab Code from cookie set in HandleLogin
	if cookie, err := r.Cookie("Code"); err == nil {
		if err = sCookie.Decode(cookie.Name, cookie.Value, &code); err != nil {
			log.Print(err)
		}
	}
	if code == "" {
		http.Redirect(w, r, "/cannotGetCookie", http.StatusTemporaryRedirect)
		return
	}

	// Exchange information with OAuth
	transport := &oauth.Transport{Config: config}
	token, err := transport.Exchange(r.FormValue("code"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	service, err := plus.New(transport.Client())
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	person, err := service.People.Get("me").Do()
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	p, err := profile.Fetch(person.Id)
	if err == profile.NotFound {
		p = &profile.Profile{
			Id:       person.Id,
			Token:    token,
			RealName: person.DisplayName,
			Url:      person.Url,
			Avatar:   person.Image.Url,
		}
	} else if err != nil {
		http.Redirect(w, r, "/profileNotFound", http.StatusTemporaryRedirect)
		return
	}

	encoded, err := sCookie.Encode("Profile", p)
	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/cannotSetCookie", http.StatusTemporaryRedirect)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "Profile",
		Value: encoded,
		Path:  "/",
	})

	http.Redirect(w, r, "/setupProfile", http.StatusTemporaryRedirect)
}
