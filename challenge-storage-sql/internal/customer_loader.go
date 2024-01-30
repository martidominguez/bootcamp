package internal

// CustomerLoader is the interface that represents the loader of customers.
type CustomerLoader interface {
	// Load loads the customers.
	Load() (customers []Customer, err error)
}
