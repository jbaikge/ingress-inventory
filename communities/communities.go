package communities

import (
	"github.com/jbaikge/ingress-inventory/faction"
)

type Community struct {
	Faction faction.Faction
	Name    string
	Url     string
}

var Communities = []Community{
	Community{
		Faction: faction.Resistance,
		Name:    "DMV Resistance",
		Url:     "https://plus.google.com/communities/103349576921336760265",
	},
	Community{
		Faction: faction.Enlightened,
		Name:    "DMV Resistance",
		Url:     "https://plus.google.com/communities/107604547228374598745",
	},
}
