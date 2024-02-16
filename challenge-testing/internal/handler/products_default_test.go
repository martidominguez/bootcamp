package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/repository"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_Get(t *testing.T) {
	t.Run("success - returns the product with the id 1", func(t *testing.T) {
		// arrange
		// - repository
		rp := repository.NewMockRepositoryProducts()
		rp.FuncSearchProducts = func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
			p = make(map[int]internal.Product)
			p[1] = internal.Product{
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "Product 1",
					Price:       100,
					SellerId:    1,
				},
			}
			return
		}
		// - handler
		hd := handler.NewProductsDefault(rp)
		hdFunc := hd.Get()
		// - request
		request := httptest.NewRequest("GET", "/product", nil)
		q := request.URL.Query()
		q.Add("id", "1")
		request.URL.RawQuery = q.Encode()
		// - response
		response := httptest.NewRecorder()

		// act
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"message": "success", "data": {"1": {"id": 1, "description": "Product 1", "price": 100, "seller_id": 1}}}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeaders, response.Header())
		require.Equal(t, 1, rp.Spy.SearchProducts)
	})
	t.Run("success - returns all the products", func(t *testing.T) {
		// arrange
		// - repository
		rp := repository.NewMockRepositoryProducts()
		rp.FuncSearchProducts = func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
			p = make(map[int]internal.Product)
			p[1] = internal.Product{
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "Product 1",
					Price:       100,
					SellerId:    1,
				},
			}
			p[2] = internal.Product{
				Id: 2,
				ProductAttributes: internal.ProductAttributes{
					Description: "Product 2",
					Price:       1000,
					SellerId:    1,
				},
			}
			return
		}
		// - handler
		hd := handler.NewProductsDefault(rp)
		hdFunc := hd.Get()
		// - request
		request := httptest.NewRequest("GET", "/product", nil)
		// - response
		response := httptest.NewRecorder()

		// act
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"message": "success", "data": {"1": {"id": 1, "description": "Product 1", "price": 100, "seller_id": 1}, "2": {"id": 2, "description": "Product 2", "price": 1000, "seller_id": 1}}}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeaders, response.Header())
		require.Equal(t, 1, rp.Spy.SearchProducts)
	})
	t.Run("error - returns the error when the id is not a number", func(t *testing.T) {
		// arrange
		// - repository
		rp := repository.NewMockRepositoryProducts()
		// - handler
		hd := handler.NewProductsDefault(rp)
		hdFunc := hd.Get()
		// - request
		request := httptest.NewRequest("GET", "/product", nil)
		q := request.URL.Query()
		q.Add("id", "not_a_number")
		request.URL.RawQuery = q.Encode()
		// - response
		response := httptest.NewRecorder()

		// act
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusBadRequest
		expectedBody := `{"message": "invalid id", "status": "Bad Request"}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeaders, response.Header())
		require.Equal(t, 0, rp.Spy.SearchProducts)
	})
	t.Run("error - returns the error when the id is not found", func(t *testing.T) {
		// arrange
		// - repository
		rp := repository.NewMockRepositoryProducts()
		rp.FuncSearchProducts = func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
			return nil, errors.New("internal server error")
		}
		// - handler
		hd := handler.NewProductsDefault(rp)
		hdFunc := hd.Get()
		// - request
		request := httptest.NewRequest("GET", "/product", nil)
		q := request.URL.Query()
		q.Add("id", "1")
		request.URL.RawQuery = q.Encode()
		// - response
		response := httptest.NewRecorder()

		// act
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusInternalServerError
		expectedBody := `{"message": "internal error", "status": "Internal Server Error"}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeaders, response.Header())
		require.Equal(t, 1, rp.Spy.SearchProducts)
	})
}
