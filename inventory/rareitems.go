package inventory

type RareItems map[Rarity]int

type Rarity int

// Rarities 1, 2 and 4 not defined yet
const (
	VeryCommon Rarity = 0
	Rare       Rarity = 3
	VeryRare   Rarity = 5
)

func NewRareItems() RareItems {
	items := make(RareItems, VeryRare+1)
	for i := Rarity(0); i < VeryRare+1; i++ {
		items[i] = 0
	}
	return items
}

func (r Rarity) Enabled() bool {
	switch r {
	case VeryCommon, Rare, VeryRare:
		return true
	default:
		return false
	}
}
