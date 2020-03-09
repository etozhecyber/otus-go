package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/etozhecyber/otus-go/calsrv/internal/domain/models"
)

type httpEvent struct {
	Events []models.Event
}

func (h *httpEvent) toJSON() ([]byte, error) {
	data, err := json.Marshal(h)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type httpError struct {
	Result string `json:"error"`
}

func (h *httpError) jsonPrint(w http.ResponseWriter, statusCode int) {
	data, err := json.Marshal(h)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}

type httpResponce struct {
	Result string `json:"result"`
}

func (h *httpResponce) jsonPrint(w http.ResponseWriter, statusCode int) {
	data, err := json.Marshal(h)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}
