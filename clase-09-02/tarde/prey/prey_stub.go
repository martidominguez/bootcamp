package prey

import "testdoubles/positioner"

// NewStub returns a new prey stub
func NewStub() *Stub {
	return &Stub{}
}

// Stub is a stub implementation of the Prey interface
type Stub struct {
	// FuncGetSpeed is a function that returns the speed of the prey (in m/s)
	FuncGetSpeed func() (speed float64)
	// FuncGetPosition is a function that returns the position of the prey
	FuncGetPosition func() (position *positioner.Position)
}

// GetSpeed returns the speed of the prey (in m/s)
func (s *Stub) GetSpeed() (speed float64) {
	speed = s.FuncGetSpeed()
	return
}

// GetPosition returns the position of the prey
func (s *Stub) GetPosition() (position *positioner.Position) {
	position = s.FuncGetPosition()
	return
}
