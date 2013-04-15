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
	hashKey = securecookie.GenerateRandomKey(64)
	// 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
	//blockKey = securecookie.GenerateRandomKey(32)
	sCookie = securecookie.New(hashKey, blockKey)
}

func main() {
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))

	log.Printf("Listening on :%s", os.Getenv("PORT"))
	panic(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
