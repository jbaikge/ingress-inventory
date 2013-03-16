package main

import (
	"fmt"
	"github.com/gorilla/context"
	"net/http"
)

func init() {
	router.HandleFunc("/", HandleIndex)
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	context.Set(r, "title", "This is a test title")
	fmt.Fprintln(w, "...")
}
