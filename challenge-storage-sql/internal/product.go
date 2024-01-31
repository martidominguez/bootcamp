package internal

// ProductAttributes is the struct that represents the attributes of a product.
type ProductAttributes struct {
	// Description is the description of the product.
	Description string
	// Price is the price of the product.
	Price float64
}

// Product is the struct that represents a product.
type Product struct {
	// Id is the unique identifier of the product.
	Id int
	// ProductAttributes is the attributes of the product.
	ProductAttributes
}

// TopProduct is the struct that represents a top product.
type TopProduct struct {
	// Description is the description of the product.
	Description string
	// Total is the amount of times the product was sold.
	Total int
}
