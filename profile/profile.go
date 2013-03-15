package profile

import (
	"github.com/jbaikge/ingress-inventory/inventory"
	"labix.org/v2/mgo/bson"
)

type Profile struct {
	Id        bson.ObjectId `bson:"_id"`
	AuthToken string
	Username  string
	inventory.Inventory
}
