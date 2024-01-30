package loader

import (
	"app/internal"
	"encoding/json"
	"os"
)

// NewCustomersJSON returns a new pointer to a CustomersJSON struct.
func NewCustomersJSON(file *os.File) *CustomersJSON {
	return &CustomersJSON{file: file}
}

// CustomersJSON is an struct that implements the LoaderCustomer interface.
type CustomersJSON struct {
	// file is the file to handle read and write operations.
	file *os.File
}

// CustomerJSON is the struct that represents a customer in the JSON file.
type CustomerJSON struct {
	Id        int    `json:"id"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Condition int    `json:"condition"`
}

// Load loads the customers from the file.
func (c *CustomersJSON) Load() (customers []internal.Customer, err error) {
	// Decode the file.
	var customersJSON []CustomerJSON
	err = json.NewDecoder(c.file).Decode(&customersJSON)
	if err != nil {
		return nil, err
	}

	// Convert the customersJSON to customers.
	for _, customerJSON := range customersJSON {
		customer := internal.Customer{
			Id: customerJSON.Id,
			CustomerAttributes: internal.CustomerAttributes{
				LastName:  customerJSON.LastName,
				FirstName: customerJSON.FirstName,
				Condition: customerJSON.Condition,
			},
		}
		customers = append(customers, customer)
	}

	return customers, nil
}
