package migrator

import "app/internal"

// NewMigratorSaleToDatabase returns a new pointer to a MigratorSaleToDatabase struct.
func NewMigratorSaleToDatabase(ld internal.SaleLoader, rp internal.RepositorySale) *MigratorSaleToDatabase {
	return &MigratorSaleToDatabase{
		ld: ld,
		rp: rp,
	}
}

// MigratorSaleToDatabase is an struct that implements the Migrator interface.
type MigratorSaleToDatabase struct {
	// ld is the loader to load the data
	ld internal.SaleLoader
	// rp is the repository to access the database
	rp internal.RepositorySale
}

// Migrate migrates the sales to the database.
func (m *MigratorSaleToDatabase) Migrate() (err error) {
	// Load the sales
	sales, err := m.ld.Load()
	if err != nil {
		return err
	}

	// Migrate the sales
	for _, sale := range sales {
		err = m.rp.Save(&sale)
		if err != nil {
			return err
		}
	}

	return nil
}
