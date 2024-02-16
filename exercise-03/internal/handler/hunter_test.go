package handler_test

import (
	"exercise-03/internal/handler"
	"exercise-03/internal/hunter"
	"exercise-03/internal/prey"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigurePrey(t *testing.T) {
	t.Run("Success: configure prey", func(t *testing.T) {
		// arrange
		// - prey: stub
		pr := prey.NewPreyStub()
		// - handler
		hd := handler.NewHunter(nil, pr)
		hdFunc := hd.ConfigurePrey()

		// act
		request := httptest.NewRequest("POST", "/", strings.NewReader(
			`{"speed": 10.0, "position": {"X": 0.0, "Y": 0.0, "Z": 0.0}}`,
		))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"message":"prey configured"}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeaders, response.Header())
	})

	t.Run("Error: invalid request body", func(t *testing.T) {
		// arrange
		// - prey: stub
		pr := prey.NewPreyStub()
		// - handler
		hd := handler.NewHunter(nil, pr)
		hdFunc := hd.ConfigurePrey()

		// act
		request := httptest.NewRequest("POST", "/", strings.NewReader(
			`{"speed": 10.0, "position": {"X": "0.0", "Y": 0.0, "Z": 0.0}}`,
		))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusBadRequest
		expectedBody := `{"message":"invalid request body"}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeaders, response.Header())
	})

}

func TestHunt(t *testing.T) {
	t.Run("Success: hunt", func(t *testing.T) {
		// arrange
		// - hunter: stub
		ht := hunter.NewHunterMock()
		ht.HuntFunc = func(pr prey.Prey) (duration float64, err error) {
			return 100.0, nil
		}
		hd := handler.NewHunter(ht, nil)
		hdFunc := hd.Hunt()
		// act
		request := httptest.NewRequest("POST", "/", nil)
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"message":"hunt done","data":{"success":true,"duration":100.0}}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		expectedCallHunt := 1
		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeaders, response.Header())
		require.Equal(t, expectedCallHunt, ht.Calls.Hunt)
	})
}
