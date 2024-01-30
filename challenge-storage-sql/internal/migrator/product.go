package migrator

import "app/internal"

func NewMigratorProductToDatabase(ld internal.ProductLoader, rp internal.RepositoryProduct) *MigratorProductToDatabase {
	return &MigratorProductToDatabase{
		ld: ld,
		rp: rp,
	}
}

// MigratorProductToDatabase is the implementation of the interface MigratorProduct.
type MigratorProductToDatabase struct {
	// ld is the loader to load the data
	ld internal.ProductLoader
	// rp is the repository to access the database
	rp internal.RepositoryProduct
}

func (m *MigratorProductToDatabase) Migrate() (err error) {
	// Load the products.
	products, err := m.ld.Load()
	if err != nil {
		return err
	}

	// Save the products in the database.
	for _, product := range products {
		err = m.rp.Save(&product)
		if err != nil {
			return err
		}
	}

	return nil
}
