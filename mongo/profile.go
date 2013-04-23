package mongo

import (
	"github.com/jbaikge/ingress-inventory/communities"
	"github.com/jbaikge/ingress-inventory/profile"
	"strings"
)

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
