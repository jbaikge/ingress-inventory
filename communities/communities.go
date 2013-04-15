package communities

import (
	"github.com/jbaikge/ingress-inventory/faction"
)

type Community struct {
	Type faction.Faction
	Name string
	Url  string
}

var Communities = []Community{
	Community{
		Type: faction.Resistance,
		Name: "DMV Resistance",
		Url: 
		},
}
