package presentation

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sdet-ozon/internal/mock_scenarios/domain"
)

type SetupManager interface {
	RegisterScenario(ctx context.Context, scenario *domain.MockScenario) error
	DeleteScenario(ctx context.Context, testID string) error
}

type RateProvider interface {
	GetExchangeRate(ctx context.Context, testID string) (*domain.MockScenario, error)
}

type Handler struct {
	setup SetupManager
	rate  RateProvider
}

func NewHandler(s SetupManager, r RateProvider) *Handler {
	return &Handler{setup: s, rate: r}
}

func (h *Handler) SetupHandler(w http.ResponseWriter, r *http.Request) {
	var dto SetupMockDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		w.Write([]byte(`{"error": "invalid json"}`))
		return
	}

	if dto.TestID == "" {
		w.Write([]byte(`{"error": "test_id is required in json"}`))
		return
	}

	scenario, err := ToDomain(dto)
	if err != nil {
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}

	if err := h.setup.RegisterScenario(r.Context(), scenario); err != nil {
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}

	w.Write([]byte(`{"success": true}`))
}

func (h *Handler) GetCbrXmlHandler(w http.ResponseWriter, r *http.Request) {
	testID := r.URL.Query().Get("test_id")

	if testID == "" {
		w.Write([]byte("Internal Server Error: test_id missing"))
		return
	}

	scenario, err := h.rate.GetExchangeRate(r.Context(), testID)
	if err != nil {
		w.Write([]byte("Internal Server Error: scenario not found"))
		return
	}

	if scenario.StatusCode == 500 {
		w.Write([]byte("Internal Server Error"))
		return
	}

	if scenario.StatusCode == 200 {
		xmlData := ToXMLResponse(scenario).ToXML()
		w.Write(xmlData)
		return
	}
}

func (h *Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	testID := r.URL.Query().Get("test_id")

	if testID == "" {
		w.Write([]byte(`{"error": "test_id is required"}`))
		return
	}

	if err := h.setup.DeleteScenario(r.Context(), testID); err != nil {
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}

	w.Write([]byte(`{"message": "scenario deleted"}`))
}
