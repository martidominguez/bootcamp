package internal

import "errors"

var (
	// ErrCustomerInternalServerError is the error when the server has an internal error.
	ErrCustomerInternalServerError = errors.New("repository: internal server error")
)

// CustomerAttributes is the struct that represents the attributes of a customer.
type CustomerAttributes struct {
	// FirstName is the first name of the customer.
	FirstName string
	// LastName is the last name of the customer.
	LastName string
	// Condition is the condition of the customer.
	Condition int
}

// Customer is the struct that represents a customer.
type Customer struct {
	// Id is the unique identifier of the customer.
	Id int
	// CustomerAttributes is the attributes of the customer.
	CustomerAttributes
}

// CustomerInvoicesByCondition is the struct that represents the total of the invoices by the customer condition.
type CustomerInvoicesByCondition struct {
	// Condition is the condition of the customer.
	Condition int
	// Total is the total of the invoices.
	Total float64
}

// TopCustomer is the struct that represents a top customer.
type TopCustomer struct {
	// FirstName is the first name of the customer.
	FirstName string
	// LastName is the last name of the customer.
	LastName string
	// Total is the total of the invoices.
	Total float64
}
