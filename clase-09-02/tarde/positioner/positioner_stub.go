package positioner

// NewStub returns a new positioner stub
func NewStub() *Stub {
	return &Stub{}
}

// Stub is a stub implementation of the Positioner interface
type Stub struct {
	// FuncGetLinearDistance is a function that returns the linear distance between 2 positions (in meters)
	FuncGetLinearDistance func(from, to *Position) (linearDistance float64)
}

// GetLinearDistance returns the linear distance between 2 positions (in meters)
func (s *Stub) GetLinearDistance(from, to *Position) (linearDistance float64) {
	linearDistance = s.FuncGetLinearDistance(from, to)
	return
}
