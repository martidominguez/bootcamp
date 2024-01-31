package internal

// ServiceCustomer is the interface that wraps the basic methods that a customer service should implement.
type ServiceCustomer interface {
	// FindAll returns all customers
	FindAll() (c []Customer, err error)
	// Save saves a customer
	Save(c *Customer) (err error)
	// FindInvoicesByCondition returns all invoices for a customer by condition
	FindInvoicesByCondition() (c []CustomerInvoicesByCondition, err error)
	// FindTop5Customers returns the top customers by total of invoices
	FindTop5Customers() (c []TopCustomer, err error)
}
