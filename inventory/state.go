package inventory

import (
	"time"
)

type State struct {
	Time       time.Time
	AP         int
	Mods       Mods
	PortalKeys int
	PowerCubes PowerCubes
	Resonators Resonators
	XMPs       XMPs
}

type PowerCubes []int

type Resonators []int

type XMPs []int

type Mods RareItems

const (
	NumMod       = VeryRare + 1
	NumPowerCube = 8
	NumResonator = 8
	NumXMP       = 8
)

var Levels = []int{1e4, 2e4, 4e4, 7e4, 15e4, 3e5, 6e5, 12e5}

func NewState() State {
	return State{
		Time:       time.Now(),
		PowerCubes: make(PowerCubes, NumPowerCube),
		Resonators: make(Resonators, NumResonator),
		Mods:       Mods(NewRareItems()),
		XMPs:       make(XMPs, NumXMP),
	}
}

func (s State) Level() int {
	for i, max := range Levels {
		if max > s.AP {
			return i + 1
		}
	}
	return 0
}
