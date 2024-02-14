package simulator

func NewSimulatorMock() *SimulatorMock {
	return &SimulatorMock{
		FuncCanCatch: func(hunter, prey *Subject) (canCatch bool) {
			return
		},
	}
}

type SimulatorMock struct {
	FuncCanCatch func(hunter, prey *Subject) (canCatch bool)

	Call struct {
		CanCatch int
	}
}

func (s *SimulatorMock) CanCatch(hunter, prey *Subject) (canCatch bool) {
	s.Call.CanCatch++
	canCatch = s.FuncCanCatch(hunter, prey)
	return
}
