package simulator

// NewMock returns a new simulator mock
func NewMock() *SimulatorMock {
	return &SimulatorMock{}
}

// CatchSimulatorMock is a mock implementation of CatchSimulator
type SimulatorMock struct {
	// FuncCanCatch is a function that returns if the hunter can catch the prey
	FuncCanCatch func(hunter, prey *Subject) (canCatch bool)
	// Spy
	Calls struct {
		// FuncCanCatch is the number of times that FuncCanCatch was called
		FuncCanCatch int
	}
}

// CanCatch returns if the hunter can catch the prey
func (s *SimulatorMock) CanCatch(hunter, prey *Subject) (canCatch bool) {
	s.Calls.FuncCanCatch++
	canCatch = s.FuncCanCatch(hunter, prey)
	return
}
