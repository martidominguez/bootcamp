package main

import (
	"app/internal/application"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// config
	cfg := &application.ApplicationMigrateConfig{
		Db: &mysql.Config{
			User:   "root",
			Passwd: "enter_your_password_here",
			Net:    "tcp",
			Addr:   "localhost:3306",
			DBName: "fantasy_products",
		},
		CustomersFile: "./docs/db/json/customers.json",
		InvoicesFile:  "./docs/db/json/invoices.json",
		ProductsFile:  "./docs/db/json/products.json",
		SalesFile:     "./docs/db/json/sales.json",
	}

	app := application.NewApplicationMigrate(cfg)
	defer app.CloseFiles()

	// set up
	err := app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}

	// run
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	// success
	fmt.Println("migration completed successfully")
}
