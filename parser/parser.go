package parser

import (
	"html/template"
	"net/http"
	"strings"
)

var cache = map[string]*template.Template{}

func Render(w http.ResponseWriter, ctx *Context, files ...string) (err error) {
	files = append(files, "base.html")
	key := strings.Join(files, "_")

	reparse, err := Reparse(files...)
	if err != nil {
		return
	}

	t, ok := cache[key]
	if !ok {
		reparse = true
	}

	if reparse {
		if t, err = template.ParseFiles(RealPaths(files)...); err != nil {
			return
		}
		cache[key] = t
	}

	return t.Execute(w, ctx)
}

func Reparse(files ...string) (reparse bool, err error) {
	var modified bool
	for _, file := range RealPaths(files) {
		f, ok := fileset[file]
		if !ok {
			f = &File{Path: file}
			fileset[file] = f
		}
		if modified, err = f.Modified(); err != nil {
			return
		}
		reparse = reparse || modified
	}
	return
}
