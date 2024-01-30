package repository

import (
	"app/internal"
	"database/sql"
	"errors"
)

// NewRepositoryProductDatabase creates a new repository for products.
func NewRepositoryProductDatabase(db *sql.DB) (r *RepositoryProductDatabase) {
	r = &RepositoryProductDatabase{db: db}
	return
}

// RepositoryProductDatabase is a repository for products.
type RepositoryProductDatabase struct {
	// db is the underlying database.
	db *sql.DB
}

func (r *RepositoryProductDatabase) FindById(id int) (p internal.Product, err error) {
	// execute query
	row := r.db.QueryRow("SELECT id, name, quantity, code_value, is_published, expiration, price FROM products WHERE id = ?", id)
	if err := row.Err(); err != nil {
		return internal.Product{}, err
	}

	// scan row
	var product internal.Product
	if err := row.Scan(&product.Id, &product.Name, &product.Quantity, &product.CodeValue, &product.IsPublished, &product.Expiration, &product.Price); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return internal.Product{}, internal.ErrRepositoryProductNotFound
		}
	}

	return product, err
}

func (r *RepositoryProductDatabase) Save(product *internal.Product) (err error) {
	// execute query
	result, err := r.db.Exec(
		"INSERT INTO products (name, quantity, code_value, is_published, expiration, price) VALUES (?, ?, ?, ?, ?, ?)",
		product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price,
	)
	if err != nil {
		return err
	}

	// get last inserted id
	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	product.Id = int(lastInsertedId)

	return err
}

func (r *RepositoryProductDatabase) UpdateOrSave(product *internal.Product) (err error) {
	// execute query
	_, err = r.db.Exec(
		"UPDATE products SET name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ? WHERE id = ?",
		product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price, product.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryProductDatabase) Update(product *internal.Product) (err error) {
	// execute query
	_, err = r.db.Exec(
		"UPDATE products SET name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ? WHERE id = ?",
		product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price, product.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryProductDatabase) Delete(id int) (err error) {
	// execute query
	_, err = r.db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
