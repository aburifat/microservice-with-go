package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/aburifat/microservice-with-go/metadata/internal/controller/metadata"
	"github.com/aburifat/microservice-with-go/metadata/internal/repository"
	"github.com/aburifat/microservice-with-go/metadata/pkg/model"
)

type Handler struct {
	controller *metadata.Controller
}

func New(controller *metadata.Controller) *Handler {
	return &Handler{controller: controller}
}

func (h *Handler) GetMetadata(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := req.Context()
	metadata, err := h.controller.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(metadata); err != nil {
		log.Printf("JSON encoding error: %v\n", err)
	}
}

func (h *Handler) PutMetadata(w http.ResponseWriter, req *http.Request) {
	var metadata model.Metadata
	if err := json.NewDecoder(req.Body).Decode(&metadata); err != nil {
		log.Printf("JSON decoding error: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := req.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := req.Context()
	if err := h.controller.Put(ctx, id, &metadata); err != nil {
		log.Printf("Controller put error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
