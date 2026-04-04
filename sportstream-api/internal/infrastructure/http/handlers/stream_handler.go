package handlers

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	streamSvc "github.com/jpsdeveloper/sportstream-api/internal/application/stream"
	"github.com/jpsdeveloper/sportstream-api/internal/domain"
	"github.com/jpsdeveloper/sportstream-api/internal/pkg/httputil"
)

type StreamHandler struct {
	service *streamSvc.Service
}

func NewStreamHandler(service *streamSvc.Service) *StreamHandler {
	return &StreamHandler{service: service}
}

func (h *StreamHandler) List(w http.ResponseWriter, r *http.Request) {
	filter := streamSvc.StreamFilter{
		Status: r.URL.Query().Get("status"),
		Type:   r.URL.Query().Get("type"),
	}

	streams, err := h.service.List(filter)
	if err != nil {
		httputil.JSONError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to list streams", err.Error())
		return
	}
	if streams == nil {
		streams = []domain.Stream{}
	}
	httputil.JSON(w, http.StatusOK, streams, len(streams))
}

func (h *StreamHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		httputil.JSONError(w, http.StatusBadRequest, "INVALID_ID", "Invalid stream ID", err.Error())
		return
	}

	stream, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, domain.ErrStreamNotFound) {
			httputil.JSONError(w, http.StatusNotFound, "NOT_FOUND", "Stream not found", "")
			return
		}
		httputil.JSONError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to get stream", err.Error())
		return
	}
	httputil.JSON(w, http.StatusOK, stream, 1)
}

func (h *StreamHandler) GetByClubID(w http.ResponseWriter, r *http.Request) {
	clubID, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		httputil.JSONError(w, http.StatusBadRequest, "INVALID_ID", "Invalid club ID", err.Error())
		return
	}

	streams, err := h.service.GetByClubID(clubID)
	if err != nil {
		if errors.Is(err, domain.ErrClubNotFound) {
			httputil.JSONError(w, http.StatusNotFound, "NOT_FOUND", "Club not found", "")
			return
		}
		httputil.JSONError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to get club streams", err.Error())
		return
	}
	if streams == nil {
		streams = []domain.Stream{}
	}
	httputil.JSON(w, http.StatusOK, streams, len(streams))
}

func (h *StreamHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input streamSvc.CreateStreamInput
	if err := httputil.DecodeJSON(r, &input); err != nil {
		httputil.JSONError(w, http.StatusBadRequest, "INVALID_BODY", "Invalid request body", err.Error())
		return
	}

	stream, err := h.service.Create(input)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidInput) {
			httputil.JSONError(w, http.StatusBadRequest, "INVALID_INPUT", "Invalid stream data", "")
			return
		}
		if errors.Is(err, domain.ErrInvalidStreamType) {
			httputil.JSONError(w, http.StatusBadRequest, "INVALID_INPUT", "Invalid stream type", "")
			return
		}
		if errors.Is(err, domain.ErrClubNotFound) {
			httputil.JSONError(w, http.StatusBadRequest, "INVALID_INPUT", "Club not found", "")
			return
		}
		httputil.JSONError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to create stream", err.Error())
		return
	}
	httputil.JSON(w, http.StatusCreated, stream, 1)
}

func (h *StreamHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		httputil.JSONError(w, http.StatusBadRequest, "INVALID_ID", "Invalid stream ID", err.Error())
		return
	}

	var input streamSvc.UpdateStreamInput
	if err := httputil.DecodeJSON(r, &input); err != nil {
		httputil.JSONError(w, http.StatusBadRequest, "INVALID_BODY", "Invalid request body", err.Error())
		return
	}

	stream, err := h.service.Update(id, input)
	if err != nil {
		if errors.Is(err, domain.ErrStreamNotFound) {
			httputil.JSONError(w, http.StatusNotFound, "NOT_FOUND", "Stream not found", "")
			return
		}
		httputil.JSONError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to update stream", err.Error())
		return
	}
	httputil.JSON(w, http.StatusOK, stream, 1)
}

func (h *StreamHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		httputil.JSONError(w, http.StatusBadRequest, "INVALID_ID", "Invalid stream ID", err.Error())
		return
	}

	var input streamSvc.UpdateStatusInput
	if err := httputil.DecodeJSON(r, &input); err != nil {
		httputil.JSONError(w, http.StatusBadRequest, "INVALID_BODY", "Invalid request body", err.Error())
		return
	}

	stream, err := h.service.UpdateStatus(id, input)
	if err != nil {
		if errors.Is(err, domain.ErrStreamNotFound) {
			httputil.JSONError(w, http.StatusNotFound, "NOT_FOUND", "Stream not found", "")
			return
		}
		if errors.Is(err, domain.ErrInvalidStreamStatus) {
			httputil.JSONError(w, http.StatusBadRequest, "INVALID_INPUT", "Invalid stream status", "")
			return
		}
		if errors.Is(err, domain.ErrInvalidTransition) {
			httputil.JSONError(w, http.StatusConflict, "INVALID_TRANSITION", "Invalid status transition", "")
			return
		}
		httputil.JSONError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to update stream status", err.Error())
		return
	}
	httputil.JSON(w, http.StatusOK, stream, 1)
}
