package parser

import (
	"github.com/jbaikge/ingress-inventory/profile"
)

type Context struct {
	Title       string
	Description string
	Profile     profile.Profile
	Javascripts []string
	Stylesheets []string
	Extra       interface{}
}

func (c *Context) AddJavascript(path string) {
	c.Javascripts = append(c.Javascripts, path)
}

func (c *Context) AddStylesheet(path string) {
	c.Stylesheets = append(c.Stylesheets, path)
}
