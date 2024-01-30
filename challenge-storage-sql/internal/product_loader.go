package internal

// ProductLoader is the interface that represents a loader of products.
type ProductLoader interface {
	// Load loads the products.
	Load() (products []Product, err error)
}
