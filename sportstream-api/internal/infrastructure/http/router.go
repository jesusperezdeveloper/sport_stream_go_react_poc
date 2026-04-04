package http

import (
	"net/http"

	"github.com/jpsdeveloper/sportstream-api/internal/infrastructure/http/handlers"
	"github.com/jpsdeveloper/sportstream-api/internal/infrastructure/http/middleware"
)

type RouterDeps struct {
	HealthHandler    *handlers.HealthHandler
	ClubHandler      *handlers.ClubHandler
	StreamHandler    *handlers.StreamHandler
	EventHandler     *handlers.EventHandler
	DashboardHandler *handlers.DashboardHandler
	AllowedOrigins   []string
}

func NewRouter(deps RouterDeps) http.Handler {
	mux := http.NewServeMux()

	// Health
	mux.HandleFunc("GET /api/v1/health", deps.HealthHandler.Health)

	// Clubs
	mux.HandleFunc("GET /api/v1/clubs", deps.ClubHandler.List)
	mux.HandleFunc("GET /api/v1/clubs/{id}", deps.ClubHandler.GetByID)
	mux.HandleFunc("POST /api/v1/clubs", deps.ClubHandler.Create)
	mux.HandleFunc("PUT /api/v1/clubs/{id}", deps.ClubHandler.Update)
	mux.HandleFunc("DELETE /api/v1/clubs/{id}", deps.ClubHandler.Delete)

	// Club streams
	mux.HandleFunc("GET /api/v1/clubs/{id}/streams", deps.StreamHandler.GetByClubID)

	// Streams
	mux.HandleFunc("GET /api/v1/streams", deps.StreamHandler.List)
	mux.HandleFunc("GET /api/v1/streams/{id}", deps.StreamHandler.GetByID)
	mux.HandleFunc("POST /api/v1/streams", deps.StreamHandler.Create)
	mux.HandleFunc("PUT /api/v1/streams/{id}", deps.StreamHandler.Update)
	mux.HandleFunc("PATCH /api/v1/streams/{id}/status", deps.StreamHandler.UpdateStatus)

	// Events
	mux.HandleFunc("GET /api/v1/events", deps.EventHandler.List)
	mux.HandleFunc("GET /api/v1/events/upcoming", deps.EventHandler.GetUpcoming)
	mux.HandleFunc("GET /api/v1/events/{id}", deps.EventHandler.GetByID)
	mux.HandleFunc("POST /api/v1/events", deps.EventHandler.Create)

	// Dashboard
	mux.HandleFunc("GET /api/v1/dashboard/summary", deps.DashboardHandler.Summary)

	// Apply middleware
	var handler http.Handler = mux
	handler = middleware.CORS(deps.AllowedOrigins)(handler)
	handler = middleware.Logging(handler)
	handler = middleware.Recovery(handler)

	return handler
}
