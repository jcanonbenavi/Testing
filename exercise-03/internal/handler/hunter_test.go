package handler_test

import (
	"exercise-03/internal/handler"
	"exercise-03/internal/prey"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigurePrey(t *testing.T) {
	t.Run("configure prey", func(t *testing.T) {
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
}
