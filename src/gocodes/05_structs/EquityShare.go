package structs

type EquityShare struct {
	Name              string
	Price             float64
	AvailableQuantity int64
}

func (e *EquityShare) incrementQuantity(quantity int64) {
	e.AvailableQuantity = e.AvailableQuantity + quantity
}

func (e *EquityShare) decrementQuantity(quantity int64) {
	e.AvailableQuantity = e.AvailableQuantity - quantity
}
