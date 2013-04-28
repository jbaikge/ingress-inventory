package profile

import (
	"code.google.com/p/goauth2/oauth"
	"encoding/gob"
	"errors"
	"github.com/jbaikge/ingress-inventory/communities"
	"github.com/jbaikge/ingress-inventory/inventory"
	"labix.org/v2/mgo/bson"
)

var NotFound = errors.New("Profile Not Found")

type Profile struct {
	Id              bson.ObjectId `bson:"_id,omitempty"`
	GoogleId        string
	Token           *oauth.Token
	DisplayName     string
	DisplayUsername string
	Username        string
	Url             string
	Avatar          string
	Communities     []communities.Community
	Inventory       inventory.Inventory
}

func init() {
	// Register with gob so we can put the profile info in the cookie
	gob.Register(&Profile{})
}
