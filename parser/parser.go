package parser

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

var (
	cache   = map[string]*template.Template{}
	funcMap = template.FuncMap{
		"add": func(a, b int) int { return a + b },
	}
)

func Render(w http.ResponseWriter, ctx *Context, files ...string) (err error) {
	files = append(files, "base.html")
	key := strings.Join(files, "_")

	paths := RealPaths(files)
	reparse, err := Reparse(paths)
	if err != nil {
		return
	}

	t, ok := cache[key]
	if !ok {
		reparse = true
	}

	if reparse {
		base := filepath.Base(files[0])
		t, err = template.New(base).Funcs(funcMap).ParseFiles(paths...)
		if err != nil {
			return
		}
		cache[key] = t
	}

	return t.Execute(w, ctx)
}

func Reparse(paths []string) (reparse bool, err error) {
	var modified bool
	for _, file := range paths {
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
