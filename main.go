package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	db "github.com/jbaikge/ingress-inventory/mongo"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	router   = mux.NewRouter()
	hashKey  []byte
	blockKey []byte
	sCookie  *securecookie.SecureCookie
	logOut   io.WriteCloser
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
	var err error

	if logOut, err = os.OpenFile("/tmp/access.log", os.O_CREATE|os.O_APPEND, 0666); err != nil {
		log.Fatal(err)
	}
	defer logOut.Close()

	if err = db.Connect(os.Getenv("DBHOST"), os.Getenv("DBNAME")); err != nil {
		log.Fatal(err)
	}

	http.Handle("/", handlers.CombinedLoggingHandler(logOut, router))
	assetHandler()

	log.Printf("Listening on :%s", os.Getenv("PORT"))
	panic(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
