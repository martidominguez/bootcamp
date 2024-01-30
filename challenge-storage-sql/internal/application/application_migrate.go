package application

import (
	"app/internal/loader"
	"app/internal/migrator"
	"app/internal/repository"
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

// ApplicationMigrateConfig is the struct that represents the application configuration for the migrate command.
type ApplicationMigrateConfig struct {
	Db            *mysql.Config
	CustomersFile string
	InvoicesFile  string
	ProductsFile  string
	SalesFile     string
}

// NewApplicationMigrate returns a new pointer to a ApplicationMigrate struct.
func NewApplicationMigrate(cfg *ApplicationMigrateConfig) *ApplicationMigrate {
	return &ApplicationMigrate{
		config: *cfg,
	}
}

// ApplicationMigrate is the struct that represents the application for the migrate command.
type ApplicationMigrate struct {
	// config is the application configuration.
	config ApplicationMigrateConfig
	// db is the database connection.
	db *sql.DB
	// files are the files to read.
	files []*os.File
	// customerMigrator is the migrator for the customers.
	customerMigrator *migrator.MigratorCustomerToDatabase
	// invoiceMigrator is the migrator for the invoices.
	invoiceMigrator *migrator.MigratorInvoiceToDatabase
	// productMigrator is the migrator for the products.
	productMigrator *migrator.MigratorProductToDatabase
	// saleMigrator is the migrator for the sales.
	saleMigrator *migrator.MigratorSaleToDatabase
}

func (a *ApplicationMigrate) CloseFiles() {
	for _, file := range a.files {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}
	return
}

func (a *ApplicationMigrate) SetUp() (err error) {
	// dependencies
	// - db connection
	a.db, err = sql.Open("mysql", a.config.Db.FormatDSN())
	if err != nil {
		return err
	}
	// - db ping
	err = a.db.Ping()
	if err != nil {
		return err
	}

	// - open files
	customerFile, err := os.Open(a.config.CustomersFile)
	if err != nil {
		return err
	}
	a.files = append(a.files, customerFile)
	invoiceFile, err := os.Open(a.config.InvoicesFile)
	if err != nil {
		return err
	}
	a.files = append(a.files, invoiceFile)
	productFile, err := os.Open(a.config.ProductsFile)
	if err != nil {
		return err
	}
	a.files = append(a.files, productFile)
	saleFile, err := os.Open(a.config.SalesFile)
	if err != nil {
		return err
	}
	a.files = append(a.files, saleFile)

	// - create migrators
	loaderCustomer := loader.NewCustomersJSON(customerFile)
	repositoryCustomer := repository.NewCustomersMySQL(a.db)
	migratorCustomer := migrator.NewMigratorCustomerToDatabase(loaderCustomer, repositoryCustomer)

	loaderInvoice := loader.NewInvoicesJSON(invoiceFile)
	repositoryInvoice := repository.NewInvoicesMySQL(a.db)
	migratorInvoice := migrator.NewMigratorInvoiceToDatabase(loaderInvoice, repositoryInvoice)

	loaderProduct := loader.NewProductsJSON(productFile)
	repositoryProduct := repository.NewProductsMySQL(a.db)
	migratorProduct := migrator.NewMigratorProductToDatabase(loaderProduct, repositoryProduct)

	loaderSale := loader.NewSalesJSON(saleFile)
	repositorySale := repository.NewSalesMySQL(a.db)
	migratorSale := migrator.NewMigratorSaleToDatabase(loaderSale, repositorySale)

	// - append the migrators
	a.customerMigrator = migratorCustomer
	a.invoiceMigrator = migratorInvoice
	a.productMigrator = migratorProduct
	a.saleMigrator = migratorSale

	return nil
}

// Run runs the application.
func (a *ApplicationMigrate) Run() (err error) {
	// Migrate the customers
	err = a.customerMigrator.Migrate()
	if err != nil {
		return err
	}

	// Migrate the invoices
	err = a.invoiceMigrator.Migrate()
	if err != nil {
		return err
	}

	// Migrate the products
	err = a.productMigrator.Migrate()
	if err != nil {
		return err
	}

	// Migrate the sales
	err = a.saleMigrator.Migrate()
	if err != nil {
		return err
	}

	return nil
}
