package main

import (
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func init() {
	for _, d := range []string{"css", "img", "js"} {
		prefix := "/" + d + "/"
		fileServer := http.FileServer(http.Dir("assets/" + d))
		handler := http.StripPrefix(prefix, fileServer)
		http.Handle(prefix, handlers.CombinedLoggingHandler(os.Stderr, handler))
	}
}
