package handler

import (
	"encoding/json"
	"errors"
	"exercise-03/internal/hunter"
	"exercise-03/internal/positioner"
	"exercise-03/internal/prey"
	"exercise-03/platform/web/response"
	"net/http"
)

// NewHunter returns a new Hunter handler.
func NewHunter(ht hunter.Hunter, pr prey.Prey) *Hunter {
	return &Hunter{ht: ht, pr: pr}
}

// Hunter returns handlers to manage hunting.
type Hunter struct {
	// ht is the Hunter interface that this handler will use
	ht hunter.Hunter
	// pr is the Prey interface that the hunter will hunt
	pr prey.Prey
}

// RequestBodyConfigPrey is an struct to configure the prey for the hunter in JSON format.
type RequestBodyConfigPrey struct {
	Speed    float64              `json:"speed"`
	Position *positioner.Position `json:"position"`
}

// ConfigurePrey configures the prey for the hunter.
func (h *Hunter) ConfigurePrey() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var requestBody RequestBodyConfigPrey
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid request body",
			})
			return
		}
		// process
		h.pr.Configure(requestBody.Speed, requestBody.Position)
		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "prey configured",
		})
	}
}

// RequestBodyConfigHunter is an struct to configure the hunter in JSON format.
type RequestBodyConfigHunter struct {
	Speed    float64              `json:"speed"`
	Position *positioner.Position `json:"position"`
}

// ConfigureHunter configures the hunter.
func (h *Hunter) ConfigureHunter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var requestBody RequestBodyConfigHunter
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid request body",
			})
			return
		}
		// process
		h.ht.Configure(requestBody.Speed, requestBody.Position)

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "hunter configured",
		})
	}
}

// Hunt hunts the prey.
func (h *Hunter) Hunt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// process
		duration, err := h.ht.Hunt(h.pr)
		if err != nil {
			if !errors.Is(err, hunter.ErrCanNotHunt) {
				response.Error(w, http.StatusInternalServerError, "internal server error")
				return
			}
		}
		ok := err == nil

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "hunt done",
			"data": map[string]any{
				"success":  ok,
				"duration": duration,
			},
		})
	}
}
