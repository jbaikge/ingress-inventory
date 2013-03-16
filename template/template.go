package template

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type manager struct {
	Template     *template.Template
	LastModified time.Time
}

var templates = map[string]*manager{}

func WriteFooter(w http.ResponseWriter, ctx *Context) {
	w.Write([]byte("FOOTER\n"))
}

func WriteHeader(w http.ResponseWriter, ctx *Context) error {
	w.Write([]byte("HEADER\n"))
	return Write(w, ctx, "header", "header.html")
}

func Write(w http.ResponseWriter, ctx *Context, name string, files ...string) error {
	w.Write([]byte("WRITER\n"))
	verifyCurrent(name, files...)
	t := templates[name]
	return t.Template.Execute(w, ctx)
}

func verifyCurrent(name string, files ...string) (err error) {
	t, ok := templates[name]
	if !ok {
		t = &manager{
			Template: template.New(name),
		}
		templates[name] = t
	}
	reparse := false
	filenames := make([]string, len(files))
	for _, file := range files {
		filename := filepath.Join("assets", "templates", file)
		in, err := os.Stat(filename)
		if err != nil {
			return err
		}
		if mtime := in.ModTime(); t.LastModified.Before(mtime) {
			reparse = true
			t.LastModified = mtime
		}
		filenames = append(filenames, filename)
	}
	if reparse {
		_, err = t.Template.ParseFiles(filenames...)
		//_, err = t.Template.Parse("<!DOCTYPE html><html><head><title>{{.Title}}</title><<body>hi")
	}
	return
}
