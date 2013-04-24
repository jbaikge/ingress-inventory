package mongo

import (
	"github.com/jbaikge/ingress-inventory/communities"
	"github.com/jbaikge/ingress-inventory/profile"
	"labix.org/v2/mgo/bson"
	"strings"
)

func FetchProfile(p *profile.Profile) (err error) {
	if !Connected() {
		return ErrNotConnected
	}

	switch {
	case p.Id != "":
		err = c.Profile.FindId(p.Id).One(p)
	case p.GoogleId != "":
		err = c.Profile.Find(bson.M{"googleid": p.GoogleId}).One(p)
	case p.Username != "":
		err = c.Profile.Find(bson.M{"username": p.Username}).One(p)
	}
	return
}

func SaveProfile(p *profile.Profile) (err error) {
	if !Connected() {
		return ErrNotConnected
	}

	// Remove unselected communities
	cNew := make([]communities.Community, 0, len(p.Communities))
	for _, com := range p.Communities {
		if com.Selected {
			cNew = append(cNew, com)
		}
	}
	p.Communities = cNew

	// Lowercase username - makes searches easier
	p.Username = strings.ToLower(p.DisplayUsername)

	if p.Id == "" {
		err = c.Profile.Insert(p)
	} else {
		err = c.Profile.Update(p.Id, p)
	}
	return
}

func UsernameRegistered(username string) (reg bool, err error) {
	if !Connected() {
		return false, ErrNotConnected
	}

	n, err := c.Profile.Find(bson.M{"username": strings.ToLower(username)}).Count()
	if err != nil {
		return
	}

	reg = n > 0
	return
}
