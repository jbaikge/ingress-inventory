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

var levels = []int{10000, 20000, 40000, 70000, 150000, 300000, 600000, 1200000}

func NewState() State {
	return State{
		Time:       time.Now(),
		Resonators: make(Resonators, MaxResonator),
		Shields:    make(Shields, MaxShield),
		XMPs:       make(XMPs, MaxXMP),
	}
}

func (s State) Level() int {
	for i, max := range levels {
		if max > s.AP {
			return i + 1
		}
	}
	return 0
}
