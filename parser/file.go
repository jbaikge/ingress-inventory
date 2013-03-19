package parser

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type File struct {
	Path         string
	LastModified time.Time
}

var (
	TemplateDir = filepath.Join("assets", "templates")
	cache       = map[string]*template.Template{}
	fileset     = map[string]*File{}
)

func (f *File) Modified() (modified bool, err error) {
	fi, err := os.Stat(f.Path)
	if err != nil {
		return false, fmt.Errorf("parser/File.Modified: %s", err)
	}
	if modTime := fi.ModTime(); !modTime.Equal(f.LastModified) {
		f.LastModified = modTime
		modified = true
	}
	return
}

func RealPaths(files []string) (paths []string) {
	paths = make([]string, len(files))
	for i, base := range files {
		paths[i] = filepath.Join(TemplateDir, base)
	}
	return
}

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
