package prey

import "exercise-01/positioner"

type PreyStub struct {
	FuncGetSpeed    func() (speed float64)
	FuncGetPosition func() (position *positioner.Position)
}

func (p *PreyStub) GetSpeed() (speed float64) {
	speed = p.FuncGetSpeed()
	return
}

// GetPosition returns the position of the prey
func (p *PreyStub) GetPosition() (position *positioner.Position) {
	position = p.FuncGetPosition()
	return
}
