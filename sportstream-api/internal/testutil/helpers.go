package testutil

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	clubSvc "github.com/jpsdeveloper/sportstream-api/internal/application/club"
	dashboardSvc "github.com/jpsdeveloper/sportstream-api/internal/application/dashboard"
	eventSvc "github.com/jpsdeveloper/sportstream-api/internal/application/event"
	streamSvc "github.com/jpsdeveloper/sportstream-api/internal/application/stream"
	apphttp "github.com/jpsdeveloper/sportstream-api/internal/infrastructure/http"
	"github.com/jpsdeveloper/sportstream-api/internal/infrastructure/http/handlers"
	"github.com/jpsdeveloper/sportstream-api/internal/infrastructure/persistence/memory"
	"github.com/jpsdeveloper/sportstream-api/internal/pkg/httputil"
)

// TestEnv holds all repos and the test server for integration tests.
type TestEnv struct {
	Server     *httptest.Server
	ClubRepo   *memory.ClubRepository
	StreamRepo *memory.StreamRepository
	EventRepo  *memory.EventRepository
}

// SetupTestRouter creates a router with in-memory repos and seed data, returns a TestEnv.
func SetupTestRouter(t *testing.T) *TestEnv {
	t.Helper()

	clubRepo := memory.NewClubRepository()
	streamRepo := memory.NewStreamRepository()
	eventRepo := memory.NewEventRepository()
	memory.SeedData(clubRepo, streamRepo, eventRepo)

	clubService := clubSvc.NewService(clubRepo)
	streamService := streamSvc.NewService(streamRepo, clubRepo)
	eventService := eventSvc.NewService(eventRepo, clubRepo)
	dashboardService := dashboardSvc.NewService(clubRepo, streamRepo, eventRepo)

	router := apphttp.NewRouter(apphttp.RouterDeps{
		HealthHandler:    handlers.NewHealthHandler("test-v1"),
		ClubHandler:      handlers.NewClubHandler(clubService),
		StreamHandler:    handlers.NewStreamHandler(streamService),
		EventHandler:     handlers.NewEventHandler(eventService),
		DashboardHandler: handlers.NewDashboardHandler(dashboardService),
		AllowedOrigins:   []string{"*"},
	})

	server := httptest.NewServer(router)
	t.Cleanup(func() { server.Close() })

	return &TestEnv{
		Server:     server,
		ClubRepo:   clubRepo,
		StreamRepo: streamRepo,
		EventRepo:  eventRepo,
	}
}

// MakeRequest sends an HTTP request and returns the response.
func MakeRequest(t *testing.T, method, url string, body interface{}) *http.Response {
	t.Helper()

	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			t.Fatalf("failed to marshal request body: %v", err)
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}
	return resp
}

// ParseSuccessResponse parses a JSON success response body.
func ParseSuccessResponse(t *testing.T, resp *http.Response) httputil.SuccessResponse {
	t.Helper()
	defer resp.Body.Close()

	var result httputil.SuccessResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode success response: %v", err)
	}
	return result
}

// ParseErrorResponse parses a JSON error response body.
func ParseErrorResponse(t *testing.T, resp *http.Response) httputil.ErrorResponse {
	t.Helper()
	defer resp.Body.Close()

	var result httputil.ErrorResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode error response: %v", err)
	}
	return result
}

// ParseDataAsSlice extracts the data field from a SuccessResponse as a slice of maps.
func ParseDataAsSlice(t *testing.T, resp httputil.SuccessResponse) []map[string]interface{} {
	t.Helper()
	raw, ok := resp.Data.([]interface{})
	if !ok {
		t.Fatalf("expected data to be a slice, got %T", resp.Data)
	}
	result := make([]map[string]interface{}, len(raw))
	for i, item := range raw {
		m, ok := item.(map[string]interface{})
		if !ok {
			t.Fatalf("expected item %d to be a map, got %T", i, item)
		}
		result[i] = m
	}
	return result
}

// ParseDataAsMap extracts the data field from a SuccessResponse as a single map.
func ParseDataAsMap(t *testing.T, resp httputil.SuccessResponse) map[string]interface{} {
	t.Helper()
	m, ok := resp.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("expected data to be a map, got %T", resp.Data)
	}
	return m
}
