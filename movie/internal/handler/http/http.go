package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/aburifat/microservice-with-go/movie/internal/controller/movie"
)

type Handler struct {
	ctrl *movie.Controller
}

func New(ctrl *movie.Controller) *Handler {
	return &Handler{
		ctrl: ctrl,
	}
}

func (h *Handler) GetMovieDetails(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	details, err := h.ctrl.Get(req.Context(), id)
	if err != nil && errors.Is(err, movie.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Response get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(details); err != nil {
		log.Printf("Response encoding error: %v\n", err)
	}
}
