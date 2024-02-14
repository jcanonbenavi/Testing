package positioner

func NewPositionerStub() *PositionerStub {
	return &PositionerStub{}
}

type PositionerStub struct {
	FuncGetLinearDistance func(from, to *Position) (linearDistance float64)
}

func (p *PositionerStub) GetLinearDistance(from, to *Position) (linearDistance float64) {
	linearDistance = p.FuncGetLinearDistance(from, to)
	return
}
