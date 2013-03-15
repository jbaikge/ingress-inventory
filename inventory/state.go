package inventory

import (
	"time"
)

type State struct {
	Time time.Time
	AP   int
	Resonators
	XMPs
	Shields
	PortalKeys
}

type PortalKeys int

type Resonators []int

type XMPs []int

type Shields []int

type ShieldType int

const (
	VeryCommon ShieldType = iota
	Rare
	VeryRare
)

const (
	MaxResonator = 8
	MaxShield    = VeryRare
	MaxXMP       = 8
)

func NewState() State {
	return Inventory{
		Time:       time.Now(),
		Resonators: make(Resonators, MaxResonator),
		Shields:    make(Shields, MaxShield),
		XMPs:       make(XMPs, MaxXMP),
	}
}
