package handlers

import (
	"net/http"

	"github.com/jpsdeveloper/sportstream-api/internal/pkg/httputil"
)

type HealthHandler struct {
	version string
}

func NewHealthHandler(version string) *HealthHandler {
	return &HealthHandler{version: version}
}

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"version": h.version,
	}
	httputil.JSON(w, http.StatusOK, data, 1)
}
