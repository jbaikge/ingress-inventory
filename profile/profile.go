package profile

import (
	"code.google.com/p/goauth2/oauth"
	"encoding/gob"
	"github.com/jbaikge/ingress-inventory/faction"
	"github.com/jbaikge/ingress-inventory/inventory"
)

type Profile struct {
	Id              string
	Token           *oauth.Token
	DisplayName     string
	DisplayUsername string
	Username        string
	Url             string
	Avatar          string
	Faction         faction.Faction
	Inventory       inventory.Inventory
}

func init() {
	// Register with gob so we can put the profile info in the cookie
	gob.Register(&Profile{})
}
