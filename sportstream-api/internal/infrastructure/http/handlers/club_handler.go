package handlers

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	clubSvc "github.com/jpsdeveloper/sportstream-api/internal/application/club"
	"github.com/jpsdeveloper/sportstream-api/internal/domain"
	"github.com/jpsdeveloper/sportstream-api/internal/pkg/httputil"
)

type ClubHandler struct {
	service *clubSvc.Service
}

func NewClubHandler(service *clubSvc.Service) *ClubHandler {
	return &ClubHandler{service: service}
}

func (h *ClubHandler) List(w http.ResponseWriter, r *http.Request) {
	clubs, err := h.service.List()
	if err != nil {
		httputil.JSONError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to list clubs", err.Error())
		return
	}
	httputil.JSON(w, http.StatusOK, clubs, len(clubs))
}

func (h *ClubHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		httputil.JSONError(w, http.StatusBadRequest, "INVALID_ID", "Invalid club ID", err.Error())
		return
	}

	club, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, domain.ErrClubNotFound) {
			httputil.JSONError(w, http.StatusNotFound, "NOT_FOUND", "Club not found", "")
			return
		}
		httputil.JSONError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to get club", err.Error())
		return
	}
	httputil.JSON(w, http.StatusOK, club, 1)
}

func (h *ClubHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input clubSvc.CreateClubInput
	if err := httputil.DecodeJSON(r, &input); err != nil {
		httputil.JSONError(w, http.StatusBadRequest, "INVALID_BODY", "Invalid request body", err.Error())
		return
	}

	club, err := h.service.Create(input)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidInput) {
			httputil.JSONError(w, http.StatusBadRequest, "INVALID_INPUT", "Name and sport are required", "")
			return
		}
		httputil.JSONError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to create club", err.Error())
		return
	}
	httputil.JSON(w, http.StatusCreated, club, 1)
}

func (h *ClubHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		httputil.JSONError(w, http.StatusBadRequest, "INVALID_ID", "Invalid club ID", err.Error())
		return
	}

	var input clubSvc.UpdateClubInput
	if err := httputil.DecodeJSON(r, &input); err != nil {
		httputil.JSONError(w, http.StatusBadRequest, "INVALID_BODY", "Invalid request body", err.Error())
		return
	}

	club, err := h.service.Update(id, input)
	if err != nil {
		if errors.Is(err, domain.ErrClubNotFound) {
			httputil.JSONError(w, http.StatusNotFound, "NOT_FOUND", "Club not found", "")
			return
		}
		httputil.JSONError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to update club", err.Error())
		return
	}
	httputil.JSON(w, http.StatusOK, club, 1)
}

func (h *ClubHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		httputil.JSONError(w, http.StatusBadRequest, "INVALID_ID", "Invalid club ID", err.Error())
		return
	}

	if err := h.service.Delete(id); err != nil {
		if errors.Is(err, domain.ErrClubNotFound) {
			httputil.JSONError(w, http.StatusNotFound, "NOT_FOUND", "Club not found", "")
			return
		}
		httputil.JSONError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to delete club", err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
