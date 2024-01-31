package repository

import (
	"database/sql"

	"app/internal"
)

// NewCustomersMySQL creates new mysql repository for customer entity.
func NewCustomersMySQL(db *sql.DB) *CustomersMySQL {
	return &CustomersMySQL{db}
}

// CustomersMySQL is the MySQL repository implementation for customer entity.
type CustomersMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all customers from the database.
func (r *CustomersMySQL) FindAll() (c []internal.Customer, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `first_name`, `last_name`, `condition` FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var cs internal.Customer
		// scan the row into the customer
		err := rows.Scan(&cs.Id, &cs.FirstName, &cs.LastName, &cs.Condition)
		if err != nil {
			return nil, err
		}
		// append the customer to the slice
		c = append(c, cs)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the customer into the database.
func (r *CustomersMySQL) Save(c *internal.Customer) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO customers (`first_name`, `last_name`, `condition`) VALUES (?, ?, ?)",
		(*c).FirstName, (*c).LastName, (*c).Condition,
	)
	if err != nil {
		return err
	}

	// get the last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set the id
	(*c).Id = int(id)

	return
}

// FindInvoicesByCondition returns the total of the invoices by the customer condition.
func (r *CustomersMySQL) FindInvoicesByCondition() (c []internal.CustomerInvoicesByCondition, err error) {
	// execute the query
	rows, err := r.db.Query(
		"SELECT c.condition, ROUND(SUM(i.total),2) AS total FROM customers c JOIN invoices i ON c.id = i.customer_id GROUP BY c.condition",
	)
	if err != nil {
		return nil, internal.ErrCustomerInternalServerError
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var ci internal.CustomerInvoicesByCondition
		// scan the row into the customer
		err := rows.Scan(&ci.Condition, &ci.Total)
		if err != nil {
			return nil, internal.ErrCustomerInternalServerError
		}
		// append the customer to the slice
		c = append(c, ci)
	}
	err = rows.Err()
	if err != nil {
		return nil, internal.ErrCustomerInternalServerError
	}

	return c, nil
}

// FindTop5Customers returns the top 5 customers by total of invoices.
func (r *CustomersMySQL) FindTop5Customers() (c []internal.TopCustomer, err error) {
	// execute the query
	rows, err := r.db.Query(
		"SELECT c.first_name, c.last_name, ROUND(SUM(i.total),2) as total FROM invoices i JOIN customers c ON i.customer_id = c.id WHERE c.condition = 1 GROUP BY c.id ORDER BY total DESC LIMIT 5",
	)
	if err != nil {
		return nil, internal.ErrCustomerInternalServerError
	}

	// iterate over the rows
	for rows.Next() {
		var tc internal.TopCustomer
		// scan the row into the customer
		err := rows.Scan(&tc.FirstName, &tc.LastName, &tc.Total)
		if err != nil {
			return nil, internal.ErrCustomerInternalServerError
		}
		// append the customer to the slice
		c = append(c, tc)
	}
	err = rows.Err()
	if err != nil {
		return nil, internal.ErrCustomerInternalServerError
	}

	return c, nil
}
