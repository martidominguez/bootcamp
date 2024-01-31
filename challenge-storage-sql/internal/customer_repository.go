package internal

// RepositoryCustomer is the interface that wraps the basic methods that a customer repository should implement.
type RepositoryCustomer interface {
	// FindAll returns all customers saved in the database.
	FindAll() (c []Customer, err error)
	// Save saves a customer into the database.
	Save(c *Customer) (err error)
	// FindInvoicesByCondition returns all invoices for a customer by condition.
	FindInvoicesByCondition() (c []CustomerInvoicesByCondition, err error)
	// FindTop5Customers returns the top customers by total of invoices.
	FindTop5Customers() (c []TopCustomer, err error)
}
