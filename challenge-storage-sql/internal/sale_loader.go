package internal

// SaleLoader is the interface that represents a loader of sales.
type SaleLoader interface {
	// Load loads the sales.
	Load() (sales []Sale, err error)
}
