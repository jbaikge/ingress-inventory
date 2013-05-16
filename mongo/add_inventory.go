package mongo

import (
	"github.com/jbaikge/ingress-inventory/inventory"
	"labix.org/v2/mgo/bson"
)

func AddInventory(id interface{}, s *inventory.State) (err error) {
	if !Connected() {
		return ErrNotConnected
	}

	return c.Profile.UpdateId(id, bson.M{"$push": bson.M{"inventory.states": s}})
}
