package internal

// InvoiceLoader is the interface that represents the loader of invoices.
type InvoiceLoader interface {
	// Load loads the invoices.
	Load() (invoices []Invoice, err error)
}
