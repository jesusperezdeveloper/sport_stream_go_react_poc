package handlers_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/jpsdeveloper/sportstream-api/internal/testutil"
)

func TestEventHandler_List(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/events", nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	result := testutil.ParseSuccessResponse(t, resp)
	data := testutil.ParseDataAsSlice(t, result)

	if len(data) != 6 {
		t.Errorf("expected 6 seeded events, got %d", len(data))
	}
	if result.Meta.Total != 6 {
		t.Errorf("expected meta.total 6, got %d", result.Meta.Total)
	}
}

func TestEventHandler_List_FilterByStatus(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/events?status=upcoming", nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	result := testutil.ParseSuccessResponse(t, resp)
	data := testutil.ParseDataAsSlice(t, result)

	// Seed data has 4 upcoming events
	if len(data) != 4 {
		t.Errorf("expected 4 upcoming events, got %d", len(data))
	}
}

func TestEventHandler_List_FilterBySport(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/events?sport=tennis", nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	result := testutil.ParseSuccessResponse(t, resp)
	data := testutil.ParseDataAsSlice(t, result)

	if len(data) != 1 {
		t.Errorf("expected 1 tennis event, got %d", len(data))
	}
}

func TestEventHandler_GetByID(t *testing.T) {
	env := testutil.SetupTestRouter(t)
	event1ID := "f1111111-1111-1111-1111-111111111111"

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/events/"+event1ID, nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	result := testutil.ParseSuccessResponse(t, resp)
	data := testutil.ParseDataAsMap(t, result)

	expected := "Lazio vs Roma — Serie A Matchday 28"
	if data["title"] != expected {
		t.Errorf("expected title %q, got %v", expected, data["title"])
	}
}

func TestEventHandler_GetByID_NotFound(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/events/00000000-0000-0000-0000-000000000000", nil)
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", resp.StatusCode)
	}

	errResp := testutil.ParseErrorResponse(t, resp)
	if errResp.Error.Code != "NOT_FOUND" {
		t.Errorf("expected error code 'NOT_FOUND', got %q", errResp.Error.Code)
	}
}

func TestEventHandler_GetByID_InvalidID(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/events/invalid", nil)
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}

func TestEventHandler_GetUpcoming(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/events/upcoming", nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	result := testutil.ParseSuccessResponse(t, resp)
	data := testutil.ParseDataAsSlice(t, result)

	// Seed data has 4 upcoming events with future start times
	if len(data) != 4 {
		t.Errorf("expected 4 upcoming events, got %d", len(data))
	}
}

func TestEventHandler_Create(t *testing.T) {
	env := testutil.SetupTestRouter(t)
	lazioID := "a1b2c3d4-e5f6-7890-abcd-ef1234567890"

	body := map[string]interface{}{
		"club_id":     lazioID,
		"title":       "New Event",
		"description": "A new event",
		"venue":       "Test Arena",
		"sport":       "football",
		"start_time":  time.Now().Add(72 * time.Hour).Format(time.RFC3339),
	}

	resp := testutil.MakeRequest(t, http.MethodPost, env.Server.URL+"/api/v1/events", body)
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}

	result := testutil.ParseSuccessResponse(t, resp)
	data := testutil.ParseDataAsMap(t, result)

	if data["title"] != "New Event" {
		t.Errorf("expected title 'New Event', got %v", data["title"])
	}
	if data["status"] != "upcoming" {
		t.Errorf("expected status 'upcoming', got %v", data["status"])
	}
}

func TestEventHandler_Create_InvalidBody_MissingTitle(t *testing.T) {
	env := testutil.SetupTestRouter(t)
	lazioID := "a1b2c3d4-e5f6-7890-abcd-ef1234567890"

	body := map[string]interface{}{
		"club_id":    lazioID,
		"sport":      "football",
		"start_time": time.Now().Add(72 * time.Hour).Format(time.RFC3339),
	}

	resp := testutil.MakeRequest(t, http.MethodPost, env.Server.URL+"/api/v1/events", body)
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}

func TestEventHandler_Create_ClubNotFound(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	body := map[string]interface{}{
		"club_id":    "00000000-0000-0000-0000-000000000000",
		"title":      "Test",
		"sport":      "football",
		"start_time": time.Now().Add(72 * time.Hour).Format(time.RFC3339),
	}

	resp := testutil.MakeRequest(t, http.MethodPost, env.Server.URL+"/api/v1/events", body)
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}
