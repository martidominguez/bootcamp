package loader

import (
	"app/internal"
	"encoding/json"
	"os"
)

// NewSalesJSON returns a new pointer to a SalesJSON struct.
func NewSalesJSON(file *os.File) *SalesJSON {
	return &SalesJSON{file: file}
}

// SalesJSON is an struct that implements the SaleLoader interface.
type SalesJSON struct {
	// file is the file to read operations.
	file *os.File
}

// SaleJSON is the struct that represents a sale in the JSON file.
type SaleJSON struct {
	Id        int `json:"id"`
	ProductId int `json:"product_id"`
	InvoiceId int `json:"invoice_id"`
	Quantity  int `json:"quantity"`
}

func (s *SalesJSON) Load() (sales []internal.Sale, err error) {
	// Decode the file
	var salesJSON []SaleJSON
	err = json.NewDecoder(s.file).Decode(&salesJSON)
	if err != nil {
		return nil, err
	}

	// Convert the salesJSON to sales.
	for _, saleJSON := range salesJSON {
		sale := internal.Sale{
			Id: saleJSON.Id,
			SaleAttributes: internal.SaleAttributes{
				ProductId: saleJSON.ProductId,
				InvoiceId: saleJSON.InvoiceId,
				Quantity:  saleJSON.Quantity,
			},
		}
		sales = append(sales, sale)
	}

	return sales, nil
}
