package loader

import (
	"app/internal"
	"encoding/json"
	"os"
)

// NewProductsJSON returns a new pointer to a ProductsJSON struct.
func NewProductsJSON(file *os.File) *ProductsJSON {
	return &ProductsJSON{file: file}
}

// ProductsJSON is an struct that implements the ProductLoader interface.
type ProductsJSON struct {
	// file is the file to read operations.
	file *os.File
}

// ProductJSON is the struct that represents a product in the JSON file.
type ProductJSON struct {
	Id          int     `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (p *ProductsJSON) Load() (products []internal.Product, err error) {
	// Decode the file
	var productsJSON []ProductJSON
	err = json.NewDecoder(p.file).Decode(&productsJSON)
	if err != nil {
		return nil, err
	}

	// Convert the productsJSON to products.
	for _, productJSON := range productsJSON {
		product := internal.Product{
			Id: productJSON.Id,
			ProductAttributes: internal.ProductAttributes{
				Description: productJSON.Description,
				Price:       productJSON.Price,
			},
		}
		products = append(products, product)
	}

	return products, nil
}
