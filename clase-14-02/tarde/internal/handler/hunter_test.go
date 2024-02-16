package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testdoubles/internal/handler"
	"testdoubles/internal/hunter"
	"testdoubles/internal/prey"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandlerNewPrey(t *testing.T) {
	t.Run("success - new prey", func(t *testing.T) {
		// arrange
		// - prey: stub
		pr := prey.NewPreyStub()
		// - handler
		hd := handler.NewHunter(nil, pr)
		hdFunc := hd.ConfigurePrey()

		// act
		request := httptest.NewRequest("POST", "/hunter/configure-prey", strings.NewReader(`{"speed": 100.0, "position": {"X": 1.0, "Y": 1.0, "Z": 0.0}}`))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"message": "prey configured}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeaders, response.Header())
	})
}

func TestHandlerNewHunter(t *testing.T) {
	t.Run("error - invalid hunter", func(t *testing.T) {
		// arrange
		// - hunter: mock
		ht := hunter.NewHunterMock()
		// - handler
		hd := handler.NewHunter(ht, nil)
		hdFunc := hd.ConfigureHunter()

		// act
		request := httptest.NewRequest("POST", "/hunter/configure-hunter", strings.NewReader(`{"speed": 100.0, "name": "hunter"`))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusBadRequest
		expectedBody := `{"message": "invalid body"}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeaders, response.Header())
	})
}

func TestHandlerHunt(t *testing.T) {
	t.Run("success - hunt", func(t *testing.T) {
		// arrange
		// - hunter: mock
		ht := hunter.NewHunterMock()
		ht.HuntFunc = func(pr prey.Prey) (duration float64, err error) {
			return 100.0, nil
		}
		// - handler
		hd := handler.NewHunter(ht, nil)
		hdFunc := hd.Hunt()

		// act
		request := httptest.NewRequest("POST", "/hunter/hunt", nil)
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"message": "hunt completed successfully", "data": {"success": true, "duration": 100.0}}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		expectedCallHunt := 1
		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeaders, response.Header())
		require.Equal(t, expectedCallHunt, ht.Calls.Hunt)
	})
}
