package template

import (
	"html/template"
	"io"
	"time"
)

type manager struct {
	Template     *template.Template
	LastModified time.Time
}

var templates = map[string]*manager{}

func WriteFooter(w io.Writer, ctx *Context) {
	w.Write([]byte("FOOTER\n"))
}

func WriteHeader(w io.Writer, ctx *Context) error {
	w.Write([]byte("HEADER\n"))
	return Write(w, ctx, "header", "header.html")
}

func Write(w io.Writer, ctx *Context, name string, files ...string) error {
	w.Write([]byte("WRITER\n"))
	verifyCurrent(name, files...)
	return templates[name].Template.Execute(w, ctx)
}

func verifyCurrent(name string, files ...string) (err error) {
	t, ok := templates[name]
	if !ok {
		t = &manager{
			Template: template.New(name),
		}
		templates[name] = t
	}
	//_, err = t.Template.ParseFiles(files...)
	_, err = t.Template.Parse("<!DOCTYPE html><html><head><title>{{.Title}}</title></head><body>hi")
	return
}
