package migrator

import "app/internal"

// NewMigratorInvoiceToDatabase is the constructor of MigratorInvoiceToDatabase.
func NewMigratorInvoiceToDatabase(ld internal.InvoiceLoader, rp internal.RepositoryInvoice) *MigratorInvoiceToDatabase {
	return &MigratorInvoiceToDatabase{
		ld: ld,
		rp: rp,
	}
}

// MigrationInvoiceToDatabase is the implementation of the interface MigratorInvoice
type MigratorInvoiceToDatabase struct {
	// ld is the loader to load the data
	ld internal.InvoiceLoader
	// rp is the repository to access the database
	rp internal.RepositoryInvoice
}

func (m *MigratorInvoiceToDatabase) Migrate() (err error) {
	// Load the invoices
	invoices, err := m.ld.Load()
	if err != nil {
		return err
	}

	// Save the invoices in the database.
	for _, invoice := range invoices {
		err = m.rp.Save(&invoice)
		if err != nil {
			return err
		}
	}

	return nil
}
