package inventory

type Inventory struct {
	States []State
}

func (i Inventory) CurrentState() State {
	if len(i.States) == 0 {
		return NewState()
	}
	return i.States[len(i.States)-1]
}
