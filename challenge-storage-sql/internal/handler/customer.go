package handler

import (
	"log"
	"net/http"

	"app/internal"
	"app/platform/web/request"
	"app/platform/web/response"
)

// NewCustomersDefault returns a new CustomersDefault
func NewCustomersDefault(sv internal.ServiceCustomer) *CustomersDefault {
	return &CustomersDefault{sv: sv}
}

// CustomersDefault is a struct that returns the customer handlers
type CustomersDefault struct {
	// sv is the customer's service
	sv internal.ServiceCustomer
}

// CustomerJSON is a struct that represents a customer in JSON format
type CustomerJSON struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Condition int    `json:"condition"`
}

// GetAll returns all customers
func (h *CustomersDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		c, err := h.sv.FindAll()
		if err != nil {
			log.Println(err)
			response.Error(w, http.StatusInternalServerError, "error getting customers")
			return
		}

		// response
		// - serialize
		csJSON := make([]CustomerJSON, len(c))
		for ix, v := range c {
			csJSON[ix] = CustomerJSON{
				Id:        v.Id,
				FirstName: v.FirstName,
				LastName:  v.LastName,
				Condition: v.Condition,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "customers found",
			"data":    csJSON,
		})
	}
}

// RequestBodyCustomer is a struct that represents the request body for a customer
type RequestBodyCustomer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Condition int    `json:"condition"`
}

// Create creates a new customer
func (h *CustomersDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - body
		var reqBody RequestBodyCustomer
		err := request.JSON(r, &reqBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "error deserializing request body")
			return
		}

		// process
		// - deserialize
		c := internal.Customer{
			CustomerAttributes: internal.CustomerAttributes{
				FirstName: reqBody.FirstName,
				LastName:  reqBody.LastName,
				Condition: reqBody.Condition,
			},
		}
		// - save
		err = h.sv.Save(&c)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error saving customer")
			return
		}

		// response
		// - serialize
		cs := CustomerJSON{
			Id:        c.Id,
			FirstName: c.FirstName,
			LastName:  c.LastName,
			Condition: c.Condition,
		}
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "customer created",
			"data":    cs,
		})
	}
}

// CustomerInvoicesByConditionJSON is a struct that represents a customer's invoices by condition in JSON format
type CustomerInvoicesByConditionJSON struct {
	Condition int     `json:"condition"`
	Total     float64 `json:"total"`
}

// GetInvoicesByCondition returns all invoices for a customer by condition
func (h *CustomersDefault) GetInvoicesByCondition() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		c, err := h.sv.FindInvoicesByCondition()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error getting invoices by condition")
			return
		}

		// response
		// - serialize
		csJSON := make([]CustomerInvoicesByConditionJSON, len(c))
		for ix, v := range c {
			csJSON[ix] = CustomerInvoicesByConditionJSON{
				Condition: v.Condition,
				Total:     v.Total,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "invoices by condition found",
			"data":    csJSON,
		})
	}
}

// TopCustomerJSON is a struct that represents a top customer in JSON format
type TopCustomerJSON struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Total     float64 `json:"total"`
}

// GetTop5Customers returns the top 5 customers by total of invoices
func (h *CustomersDefault) GetTop5Customers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		c, err := h.sv.FindTop5Customers()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error getting top 5 customers")
			return
		}

		// response
		// - serialize
		csJSON := make([]TopCustomerJSON, len(c))
		for ix, v := range c {
			csJSON[ix] = TopCustomerJSON{
				FirstName: v.FirstName,
				LastName:  v.LastName,
				Total:     v.Total,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "top 5 customers found",
			"data":    csJSON,
		})
	}
}
