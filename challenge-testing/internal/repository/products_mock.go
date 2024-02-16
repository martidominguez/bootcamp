package repository

import "app/internal"

// NewMockRepositoryProducts returns a new MockRepositoryProducts.
func NewMockRepositoryProducts() *MockRepositoryProducts {
	return &MockRepositoryProducts{}
}

// MockRepositoryProducts is an struct that implements the RepositoryProducts interface.
type MockRepositoryProducts struct {
	// FuncSearchProducts is the function that imitates the SearchProducts method.
	FuncSearchProducts func(query internal.ProductQuery) (p map[int]internal.Product, err error)
	// Spy
	Spy struct {
		// SearchProducts is the number of calls to the SearchProducts method.
		SearchProducts int
	}
}

// SearchProducts returns a list of products that match the query.
func (r *MockRepositoryProducts) SearchProducts(query internal.ProductQuery) (p map[int]internal.Product, err error) {
	// Increment the number of calls to the SearchProducts method.
	r.Spy.SearchProducts++
	// Call the function that imitates the SearchProducts method.
	return r.FuncSearchProducts(query)
}
