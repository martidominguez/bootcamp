package service

import "app/internal"

// NewCustomersDefault creates new default service for customer entity.
func NewCustomersDefault(rp internal.RepositoryCustomer) *CustomersDefault {
	return &CustomersDefault{rp}
}

// CustomersDefault is the default service implementation for customer entity.
type CustomersDefault struct {
	// rp is the repository for customer entity.
	rp internal.RepositoryCustomer
}

// FindAll returns all customers.
func (s *CustomersDefault) FindAll() (c []internal.Customer, err error) {
	c, err = s.rp.FindAll()
	return
}

// Save saves the customer.
func (s *CustomersDefault) Save(c *internal.Customer) (err error) {
	err = s.rp.Save(c)
	return
}

// FindInvoicesByCondition returns all invoices for a customer by condition.
func (s *CustomersDefault) FindInvoicesByCondition() (c []internal.CustomerInvoicesByCondition, err error) {
	c, err = s.rp.FindInvoicesByCondition()
	return
}

// FindTop5Customers returns the top customers by total of invoices.
func (s *CustomersDefault) FindTop5Customers() (c []internal.TopCustomer, err error) {
	c, err = s.rp.FindTop5Customers()
	return
}
