package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"log"
	"net/http"
	"os"
)

var (
	router   = mux.NewRouter()
	hashKey  []byte
	blockKey []byte
	sCookie  *securecookie.SecureCookie
)

func init() {
	log.SetFlags(log.Lshortfile)
	hashKey = []byte(os.Getenv("HASHKEY"))
	if bKey := os.Getenv("BLOCKKEY"); bKey != "" {
		blockKey = []byte(bKey)
	}
	sCookie = securecookie.New(hashKey, blockKey)
}

func main() {
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))

	log.Printf("Listening on :%s", os.Getenv("PORT"))
	panic(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
