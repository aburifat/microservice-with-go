package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/aburifat/microservice-with-go/rating/internal/controller/rating"
	"github.com/aburifat/microservice-with-go/rating/pkg/model"
)

type Handler struct {
	ctrl *rating.Controller
}

func New(ctrl *rating.Controller) *Handler {
	return &Handler{
		ctrl: ctrl,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, req *http.Request) {
	recordID := model.RecordID(req.FormValue("id"))

	if recordID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	recordType := model.RecordType(req.FormValue("type"))
	if recordType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch req.Method {
	case http.MethodGet:
		v, err := h.ctrl.GetAggregatedRating(req.Context(), recordID, recordType)
		if err != nil && errors.Is(err, rating.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Printf("Response encoding error: %v\n", err)
		}
	case http.MethodPut:
		userID := model.UserID(req.FormValue("userId"))
		value, err := strconv.ParseFloat(req.FormValue("value"), 64)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := h.ctrl.PutRating(req.Context(), recordID, recordType, &model.Rating{
			ID:         "",
			RecordID:   recordID,
			RecordType: recordType,
			UserID:     userID,
			Value:      value,
		}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
