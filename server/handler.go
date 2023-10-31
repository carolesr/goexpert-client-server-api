package main

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *Service
}

func (h *Handler) ExchangeRateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	exchangeRate, err := h.service.GetExchangeRate(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exchangeRate)
}
