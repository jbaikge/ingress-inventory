package profile

import (
	"code.google.com/p/goauth2/oauth"
	"encoding/gob"
	"github.com/jbaikge/ingress-inventory/inventory"
)

type Profile struct {
	Id        string
	Token     *oauth.Token
	RealName  string
	Username  string
	Url       string
	Avatar    string
	Faction   Faction
	Inventory inventory.Inventory
}

type Faction string

const (
	Resistance  Faction = "RESISTANCE"
	Enlightened Faction = "ENLIGHTENED"
)

func init() {
	// Register with gob so we can put the profile info in the cookie
	gob.Register(&Profile{})
}
