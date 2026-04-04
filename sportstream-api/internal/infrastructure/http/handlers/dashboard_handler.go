package handlers

import (
	"net/http"

	dashboardSvc "github.com/jpsdeveloper/sportstream-api/internal/application/dashboard"
	"github.com/jpsdeveloper/sportstream-api/internal/pkg/httputil"
)

type DashboardHandler struct {
	service *dashboardSvc.Service
}

func NewDashboardHandler(service *dashboardSvc.Service) *DashboardHandler {
	return &DashboardHandler{service: service}
}

func (h *DashboardHandler) Summary(w http.ResponseWriter, r *http.Request) {
	summary, err := h.service.GetSummary()
	if err != nil {
		httputil.JSONError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to get dashboard summary", err.Error())
		return
	}
	httputil.JSON(w, http.StatusOK, summary, 1)
}
