package loader

import (
	"app/internal"
	"encoding/json"
	"os"
)

// NewInvoicesJSON returns a new pointer to a InvoicesJSON struct.
func NewInvoicesJSON(file *os.File) *InvoicesJSON {
	return &InvoicesJSON{file: file}
}

// InvoicesJSON is an struct that implements the LoaderInvoice interface.
type InvoicesJSON struct {
	// file is the file to read operations.
	file *os.File
}

// InvoiceJSON is the struct that represents an invoice in the JSON file.
type InvoiceJSON struct {
	Id         int     `json:"id"`
	Datetime   string  `json:"datetime"`
	CustomerId int     `json:"customer_id"`
	Total      float64 `json:"total"`
}

func (i *InvoicesJSON) Load() (invoices []internal.Invoice, err error) {
	// Decode the file
	var invoicesJSON []InvoiceJSON
	err = json.NewDecoder(i.file).Decode(&invoicesJSON)
	if err != nil {
		return nil, err
	}

	// Convert the invoicesJSON to invoices.
	for _, invoiceJSON := range invoicesJSON {
		invoice := internal.Invoice{
			Id: invoiceJSON.Id,
			InvoiceAttributes: internal.InvoiceAttributes{
				Datetime:   invoiceJSON.Datetime,
				CustomerId: invoiceJSON.CustomerId,
				Total:      invoiceJSON.Total,
			},
		}
		invoices = append(invoices, invoice)
	}

	return invoices, nil
}
