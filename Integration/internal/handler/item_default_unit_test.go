package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/integration/internal"
	"testing/integration/internal/handler"
	"testing/integration/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestFindById(t *testing.T) {
	t.Run("should return an item by its ID", func(t *testing.T) {
		// arrange
		// - Service
		sv := service.NewItemDefaultMock()
		// - Handler
		hd := handler.NewItemDefault(sv)
		// -Set the find by id method to return an item
		sv.On("FindById", 1).Return(internal.Item{
			ID:          1,
			Name:        "item 01",
			Description: "description 01",
			Price:       100.00,
		}, nil)
		// - Request
		req := httptest.NewRequest(http.MethodGet, "/items/1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - Response
		res := httptest.NewRecorder()
		// act
		hd.FindById()(res, req)
		// assert
		expectedBody := `{"data":{"id":1,"name":"item 01","description":"description 01","price":100},"message":"item found"}`

		//require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, expectedBody, res.Body.String())
	})
}
