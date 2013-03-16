package template

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const TemplateDir = filepath.Join("assets", "templates")

type File struct {
	Path         string
	LastModified time.Time
}

type FileManager struct {
	Files []File
}

func (f File) Modified() (modified bool, err error) {
	fi, err := os.Stat(f.Path)
	if err != nil {
		return false, fmt.Errorf("template/File.Modified: %s", err)
	}
	if modTime := fi.ModTime(); !modTime.Equal(f.LastModified) {
		f.LastModified = modTime
		modified = true
	}
	return
}

var templates = map[string]*manager{}

func WriteFooter(w http.ResponseWriter, ctx *Context) error {
	w.Write([]byte("FOOTER\n"))
	return nil
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
	log.Printf("Verifying %s %v", name, files)
	t, ok := templates[name]
	if !ok {
		log.Print("Building new manager")
		t = &manager{
			Template: template.New(name),
		}
		templates[name] = t
	}
	reparse := false
	filenames := make([]string, len(files))
	for i, file := range files {
		filename := filepath.Join("assets", "templates", file)
		in, err := os.Stat(filename)
		if err != nil {
			log.Printf("Could not stat %s", filename)
			return err
		}
		if mtime := in.ModTime(); t.LastModified.Before(mtime) {
			log.Printf("Last modified [%s]: %s", filename, mtime)
			reparse = true
			t.LastModified = mtime
		}
		filenames[i] = filename
	}
	log.Printf("Filenames: %+q", filenames)
	if reparse {
		log.Print("Templates changed, reparsing")
		_, err = t.Template.ParseFiles(filenames...)
		if err != nil {
			log.Printf("Error reparsing: %s", err)
		}
	}
	return
}
