package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	clubSvc "github.com/jpsdeveloper/sportstream-api/internal/application/club"
	dashboardSvc "github.com/jpsdeveloper/sportstream-api/internal/application/dashboard"
	eventSvc "github.com/jpsdeveloper/sportstream-api/internal/application/event"
	streamSvc "github.com/jpsdeveloper/sportstream-api/internal/application/stream"
	"github.com/jpsdeveloper/sportstream-api/internal/infrastructure/config"
	router "github.com/jpsdeveloper/sportstream-api/internal/infrastructure/http"
	"github.com/jpsdeveloper/sportstream-api/internal/infrastructure/http/handlers"
	"github.com/jpsdeveloper/sportstream-api/internal/infrastructure/persistence/memory"
)

func main() {
	cfg := config.Load()

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})))

	// Repositories
	clubRepo := memory.NewClubRepository()
	streamRepo := memory.NewStreamRepository()
	eventRepo := memory.NewEventRepository()

	// Seed data
	memory.SeedData(clubRepo, streamRepo, eventRepo)

	// Services
	clubService := clubSvc.NewService(clubRepo)
	streamService := streamSvc.NewService(streamRepo, clubRepo)
	eventService := eventSvc.NewService(eventRepo, clubRepo)
	dashboardService := dashboardSvc.NewService(clubRepo, streamRepo, eventRepo)

	// Handlers
	healthHandler := handlers.NewHealthHandler(cfg.Version)
	clubHandler := handlers.NewClubHandler(clubService)
	streamHandler := handlers.NewStreamHandler(streamService)
	eventHandler := handlers.NewEventHandler(eventService)
	dashboardHandler := handlers.NewDashboardHandler(dashboardService)

	// Router
	handler := router.NewRouter(router.RouterDeps{
		HealthHandler:    healthHandler,
		ClubHandler:      clubHandler,
		StreamHandler:    streamHandler,
		EventHandler:     eventHandler,
		DashboardHandler: dashboardHandler,
		AllowedOrigins:   cfg.CORSAllowedOrigins,
	})

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown
	go func() {
		slog.Info("server starting", "port", cfg.Port, "env", cfg.Env, "version", cfg.Version)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server failed", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("server shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("server forced shutdown", "error", err)
	}

	slog.Info("server stopped")
}
