package template

import (
	"io"
)

func WriteFooter(w io.Writer, ctx *Context) {
	w.Write([]byte("FOOTER\n"))
}

func WriteHeader(w io.Writer, ctx *Context) {
	w.Write([]byte("HEADER\n"))

}

func verifyCurrent()
