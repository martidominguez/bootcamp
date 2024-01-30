package migrator

import "app/internal"

// NewMigratorCustomerToDatabase is the constructor of MigratorCustomerToDatabase.
func NewMigratorCustomerToDatabase(ld internal.CustomerLoader, rp internal.RepositoryCustomer) *MigratorCustomerToDatabase {
	return &MigratorCustomerToDatabase{
		ld: ld,
		rp: rp,
	}
}

// MigratorCustomerToDatabase is the implementation of the interface MigratorCustomer
type MigratorCustomerToDatabase struct {
	// ld is the loader to load the data
	ld internal.CustomerLoader
	// rp is the repository to access the database
	rp internal.RepositoryCustomer
}

func (m *MigratorCustomerToDatabase) Migrate() (err error) {
	// Load the customers.
	customers, err := m.ld.Load()
	if err != nil {
		return err
	}

	// Save the customers in the database.
	for _, customer := range customers {
		err = m.rp.Save(&customer)
		if err != nil {
			return err
		}
	}

	return nil
}
