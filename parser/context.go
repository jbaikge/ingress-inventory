package parser

import (
	"github.com/jbaikge/ingress-inventory/profile"
)

type Context struct {
	Title       string
	Description string
	Profile     profile.Profile
}
