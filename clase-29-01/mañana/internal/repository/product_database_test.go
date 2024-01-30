package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

func init() {
	// database config
	cfg := mysql.Config{
		User:      "root",
		Passwd:    "",
		Net:       "tcp",
		Addr:      "localhost:3306",
		DBName:    "my_db",
		ParseTime: true,
	}
	// register txdb driver
	txdb.Register("txdb", "mysql", cfg.FormatDSN())
}

func TestProductMySQL_FindById(t *testing.T) {
	t.Run("sucess 01 - product found", func(t *testing.T) {
		// arrange
		// - db
		db, err := sql.Open("txdb", "test_product_find_by_id_success_01")
		require.NoError(t, err)
		defer db.Close()
		defer func(db *sql.DB) {
			// clean data from products table
			_, err := db.Exec("DELETE FROM products")
			require.NoError(t, err)
			// reset auto increment
			_, err = db.Exec("ALTER TABLE products AUTO_INCREMENT = 1")
			require.NoError(t, err)
		}(db)
		// - set up
		err = func(db *sql.DB) error {
			_, err := db.Exec(
				"INSERT INTO products (id, name, quantity, code_value, is_published, expiration, price, id_warehouse) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
				1, "product 1", 10, "code 1", true, "2021-12-31", 100.0, 1,
			)
			return err
		}(db)
		require.NoError(t, err)
		// - repository
		repository := repository.NewRepositoryProductDatabase(db)

		// act
		product, err := repository.FindById(1)

		// assert
		expectedProduct := internal.Product{
			Id: 1,
			ProductAttributes: internal.ProductAttributes{
				Name:        "product 1",
				Quantity:    10,
				CodeValue:   "code 1",
				IsPublished: true,
				Expiration: func() time.Time {
					t, _ := time.Parse("2006-01-02", "2021-12-31")
					return t
				}(),
				Price: 100.0,
			},
		}
		require.NoError(t, err)
		require.Equal(t, expectedProduct, product)
	})

	t.Run("error 01 - product not found", func(t *testing.T) {
		// arrange
		// - db
		db, err := sql.Open("txdb", "test_product_find_by_id_success_01")
		require.NoError(t, err)
		defer db.Close()
		// - repository
		repository := repository.NewRepositoryProductDatabase(db)

		// act
		product, err := repository.FindById(1)

		// asert
		expectedProduct := internal.Product{}
		expectedError := internal.ErrRepositoryProductNotFound
		require.Equal(t, expectedProduct, product)
		require.ErrorIs(t, err, expectedError)
		require.EqualError(t, err, expectedError.Error())
	})
}
