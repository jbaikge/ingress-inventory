package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type File struct {
	Path         string
	LastModified time.Time
}

var (
	TemplateDir = filepath.Join("assets", "templates")
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
