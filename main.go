package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var (
	router = mux.NewRouter()
)

func main() {
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))

	log.Printf("Listening on :%s", os.Getenv("PORT"))
	panic(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
