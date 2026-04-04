package handlers

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	eventSvc "github.com/jpsdeveloper/sportstream-api/internal/application/event"
	"github.com/jpsdeveloper/sportstream-api/internal/domain"
	"github.com/jpsdeveloper/sportstream-api/internal/pkg/httputil"
)

type EventHandler struct {
	service *eventSvc.Service
}

func NewEventHandler(service *eventSvc.Service) *EventHandler {
	return &EventHandler{service: service}
}

func (h *EventHandler) List(w http.ResponseWriter, r *http.Request) {
	filter := eventSvc.EventFilter{
		Status: r.URL.Query().Get("status"),
		Sport:  r.URL.Query().Get("sport"),
	}

	events, err := h.service.List(filter)
	if err != nil {
		httputil.JSONError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to list events", err.Error())
		return
	}
	if events == nil {
		events = []domain.Event{}
	}
	httputil.JSON(w, http.StatusOK, events, len(events))
}

func (h *EventHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		httputil.JSONError(w, http.StatusBadRequest, "INVALID_ID", "Invalid event ID", err.Error())
		return
	}

	event, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, domain.ErrEventNotFound) {
			httputil.JSONError(w, http.StatusNotFound, "NOT_FOUND", "Event not found", "")
			return
		}
		httputil.JSONError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to get event", err.Error())
		return
	}
	httputil.JSON(w, http.StatusOK, event, 1)
}

func (h *EventHandler) GetUpcoming(w http.ResponseWriter, r *http.Request) {
	events, err := h.service.GetUpcoming()
	if err != nil {
		httputil.JSONError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to get upcoming events", err.Error())
		return
	}
	if events == nil {
		events = []domain.Event{}
	}
	httputil.JSON(w, http.StatusOK, events, len(events))
}

func (h *EventHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input eventSvc.CreateEventInput
	if err := httputil.DecodeJSON(r, &input); err != nil {
		httputil.JSONError(w, http.StatusBadRequest, "INVALID_BODY", "Invalid request body", err.Error())
		return
	}

	event, err := h.service.Create(input)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidInput) {
			httputil.JSONError(w, http.StatusBadRequest, "INVALID_INPUT", "Invalid event data", "")
			return
		}
		if errors.Is(err, domain.ErrClubNotFound) {
			httputil.JSONError(w, http.StatusBadRequest, "INVALID_INPUT", "Club not found", "")
			return
		}
		httputil.JSONError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to create event", err.Error())
		return
	}
	httputil.JSON(w, http.StatusCreated, event, 1)
}
