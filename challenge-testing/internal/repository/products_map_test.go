package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepositorySearchProducst(t *testing.T) {
	t.Run("success - returns the product with the id 1", func(t *testing.T) {
		// arrange
		// - db
		db := make(map[int]internal.Product)
		db[1] = internal.Product{
			Id: 1,
			ProductAttributes: internal.ProductAttributes{
				Description: "Product 1",
				Price:       100,
				SellerId:    1,
			},
		}
		db[2] = internal.Product{
			Id: 2,
			ProductAttributes: internal.ProductAttributes{
				Description: "Product 2",
				Price:       1000,
				SellerId:    1,
			},
		}
		// - repository
		repository := repository.NewProductsMap(db)
		// - query
		query := internal.ProductQuery{
			Id: 1,
		}

		// act
		p, err := repository.SearchProducts(query)

		// assert
		productExpected := internal.Product{
			Id: 1,
			ProductAttributes: internal.ProductAttributes{
				Description: "Product 1",
				Price:       100,
				SellerId:    1,
			},
		}
		require.NoError(t, err)
		require.Equal(t, 1, len(p))
		require.Equal(t, productExpected, p[1])
	})
	t.Run("success - returns all the products", func(t *testing.T) {
		// arrange
		// - db
		db := make(map[int]internal.Product)
		db[1] = internal.Product{
			Id: 1,
			ProductAttributes: internal.ProductAttributes{
				Description: "Product 1",
				Price:       100,
				SellerId:    1,
			},
		}
		db[2] = internal.Product{
			Id: 2,
			ProductAttributes: internal.ProductAttributes{
				Description: "Product 2",
				Price:       200,
				SellerId:    2,
			},
		}
		// - repository
		repository := repository.NewProductsMap(db)
		// - query
		query := internal.ProductQuery{}

		// act
		p, err := repository.SearchProducts(query)

		// assert
		expectedResponse := make(map[int]internal.Product)
		expectedResponse[1] = internal.Product{
			Id: 1,
			ProductAttributes: internal.ProductAttributes{
				Description: "Product 1",
				Price:       100,
				SellerId:    1,
			},
		}
		expectedResponse[2] = internal.Product{
			Id: 2,
			ProductAttributes: internal.ProductAttributes{
				Description: "Product 2",
				Price:       200,
				SellerId:    2,
			},
		}
		require.NoError(t, err)
		require.Equal(t, 2, len(p))
		require.Equal(t, expectedResponse, p)
	})
	t.Run("error - negative id", func(t *testing.T) {
		// arrange
		// - db
		db := make(map[int]internal.Product)
		db[1] = internal.Product{
			Id: 1,
			ProductAttributes: internal.ProductAttributes{
				Description: "Product 1",
				Price:       100,
				SellerId:    1,
			},
		}
		// - repository
		repository := repository.NewProductsMap(db)
		// - query
		query := internal.ProductQuery{
			Id: -1,
		}

		// act
		p, err := repository.SearchProducts(query)

		// assert
		productExpected := internal.Product{
			Id: 1,
			ProductAttributes: internal.ProductAttributes{
				Description: "Product 1",
				Price:       100,
				SellerId:    1,
			},
		}
		require.NoError(t, err)
		require.Equal(t, 1, len(p))
		require.Equal(t, productExpected, p[1])
	})
}
