package main

import (
	//"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

var (
	router = mux.NewRouter()
)

func main() {
	//http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))
	http.Handle("/", router)

	panic(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
