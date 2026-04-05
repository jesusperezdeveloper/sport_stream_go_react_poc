package handlers_test

import (
	"net/http"
	"testing"

	"github.com/jpsdeveloper/sportstream-api/internal/testutil"
)

func TestStreamHandler_List(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/streams", nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	result := testutil.ParseSuccessResponse(t, resp)
	data := testutil.ParseDataAsSlice(t, result)

	if len(data) != 10 {
		t.Errorf("expected 10 seeded streams, got %d", len(data))
	}
	if result.Meta.Total != 10 {
		t.Errorf("expected meta.total 10, got %d", result.Meta.Total)
	}
}

func TestStreamHandler_List_FilterByStatus(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/streams?status=live", nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	result := testutil.ParseSuccessResponse(t, resp)
	data := testutil.ParseDataAsSlice(t, result)

	// Seed data has 2 live streams (stream1 and stream5)
	if len(data) != 2 {
		t.Errorf("expected 2 live streams, got %d", len(data))
	}

	// Verify all returned streams are live
	for i, s := range data {
		if s["status"] != "live" {
			t.Errorf("stream %d: expected status 'live', got %v", i, s["status"])
		}
	}
}

func TestStreamHandler_List_FilterByType(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/streams?type=highlight", nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	result := testutil.ParseSuccessResponse(t, resp)
	data := testutil.ParseDataAsSlice(t, result)

	// Seed data has 3 highlight streams (stream4, stream6, stream10)
	if len(data) != 3 {
		t.Errorf("expected 3 highlight streams, got %d", len(data))
	}
}

func TestStreamHandler_List_FilterByStatusNoResults(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/streams?status=nonexistent", nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	result := testutil.ParseSuccessResponse(t, resp)
	data := testutil.ParseDataAsSlice(t, result)

	if len(data) != 0 {
		t.Errorf("expected 0 streams, got %d", len(data))
	}
}

func TestStreamHandler_GetByID(t *testing.T) {
	env := testutil.SetupTestRouter(t)
	stream1ID := "11111111-1111-1111-1111-111111111111"

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/streams/"+stream1ID, nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	result := testutil.ParseSuccessResponse(t, resp)
	data := testutil.ParseDataAsMap(t, result)

	if data["title"] != "Lazio vs Roma — Serie A Matchday 28" {
		t.Errorf("unexpected title: %v", data["title"])
	}
	if data["status"] != "live" {
		t.Errorf("expected status 'live', got %v", data["status"])
	}
}

func TestStreamHandler_GetByID_NotFound(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/streams/00000000-0000-0000-0000-000000000000", nil)
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", resp.StatusCode)
	}

	errResp := testutil.ParseErrorResponse(t, resp)
	if errResp.Error.Code != "NOT_FOUND" {
		t.Errorf("expected error code 'NOT_FOUND', got %q", errResp.Error.Code)
	}
}

func TestStreamHandler_GetByID_InvalidID(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/streams/bad-id", nil)
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}

func TestStreamHandler_GetByClubID(t *testing.T) {
	env := testutil.SetupTestRouter(t)
	lazioID := "a1b2c3d4-e5f6-7890-abcd-ef1234567890"

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/clubs/"+lazioID+"/streams", nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	result := testutil.ParseSuccessResponse(t, resp)
	data := testutil.ParseDataAsSlice(t, result)

	// Lazio has 2 streams in seed data
	if len(data) != 2 {
		t.Errorf("expected 2 streams for Lazio, got %d", len(data))
	}
}

func TestStreamHandler_GetByClubID_NotFound(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/clubs/00000000-0000-0000-0000-000000000000/streams", nil)
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", resp.StatusCode)
	}
}

func TestStreamHandler_Create(t *testing.T) {
	env := testutil.SetupTestRouter(t)
	lazioID := "a1b2c3d4-e5f6-7890-abcd-ef1234567890"

	body := map[string]interface{}{
		"club_id":     lazioID,
		"title":       "New Lazio Stream",
		"description": "A new stream",
		"type":        "vod",
		"stream_url":  "https://example.com/stream.m3u8",
		"tags":        []string{"football", "test"},
	}

	resp := testutil.MakeRequest(t, http.MethodPost, env.Server.URL+"/api/v1/streams", body)
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}

	result := testutil.ParseSuccessResponse(t, resp)
	data := testutil.ParseDataAsMap(t, result)

	if data["title"] != "New Lazio Stream" {
		t.Errorf("expected title 'New Lazio Stream', got %v", data["title"])
	}
	if data["status"] != "scheduled" {
		t.Errorf("expected status 'scheduled', got %v", data["status"])
	}
}

func TestStreamHandler_Create_InvalidBody_MissingTitle(t *testing.T) {
	env := testutil.SetupTestRouter(t)
	lazioID := "a1b2c3d4-e5f6-7890-abcd-ef1234567890"

	body := map[string]interface{}{
		"club_id": lazioID,
		"type":    "live",
	}

	resp := testutil.MakeRequest(t, http.MethodPost, env.Server.URL+"/api/v1/streams", body)
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}

func TestStreamHandler_Create_InvalidType(t *testing.T) {
	env := testutil.SetupTestRouter(t)
	lazioID := "a1b2c3d4-e5f6-7890-abcd-ef1234567890"

	body := map[string]interface{}{
		"club_id": lazioID,
		"title":   "Test",
		"type":    "podcast",
	}

	resp := testutil.MakeRequest(t, http.MethodPost, env.Server.URL+"/api/v1/streams", body)
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}

func TestStreamHandler_Create_ClubNotFound(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	body := map[string]interface{}{
		"club_id": "00000000-0000-0000-0000-000000000000",
		"title":   "Test",
		"type":    "live",
	}

	resp := testutil.MakeRequest(t, http.MethodPost, env.Server.URL+"/api/v1/streams", body)
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}

func TestStreamHandler_Update(t *testing.T) {
	env := testutil.SetupTestRouter(t)
	stream1ID := "11111111-1111-1111-1111-111111111111"

	body := map[string]interface{}{
		"title": "Updated Stream Title",
	}

	resp := testutil.MakeRequest(t, http.MethodPut, env.Server.URL+"/api/v1/streams/"+stream1ID, body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	result := testutil.ParseSuccessResponse(t, resp)
	data := testutil.ParseDataAsMap(t, result)

	if data["title"] != "Updated Stream Title" {
		t.Errorf("expected title 'Updated Stream Title', got %v", data["title"])
	}
}

func TestStreamHandler_Update_NotFound(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	body := map[string]interface{}{"title": "Whatever"}
	resp := testutil.MakeRequest(t, http.MethodPut, env.Server.URL+"/api/v1/streams/00000000-0000-0000-0000-000000000000", body)
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", resp.StatusCode)
	}
}

func TestStreamHandler_UpdateStatus(t *testing.T) {
	env := testutil.SetupTestRouter(t)
	// stream3 is "scheduled", transition to "live" is valid
	stream3ID := "33333333-3333-3333-3333-333333333333"

	body := map[string]interface{}{
		"status": "live",
	}

	resp := testutil.MakeRequest(t, http.MethodPatch, env.Server.URL+"/api/v1/streams/"+stream3ID+"/status", body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	result := testutil.ParseSuccessResponse(t, resp)
	data := testutil.ParseDataAsMap(t, result)

	if data["status"] != "live" {
		t.Errorf("expected status 'live', got %v", data["status"])
	}
}

func TestStreamHandler_UpdateStatus_InvalidStatus(t *testing.T) {
	env := testutil.SetupTestRouter(t)
	stream3ID := "33333333-3333-3333-3333-333333333333"

	body := map[string]interface{}{
		"status": "bogus",
	}

	resp := testutil.MakeRequest(t, http.MethodPatch, env.Server.URL+"/api/v1/streams/"+stream3ID+"/status", body)
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}

func TestStreamHandler_UpdateStatus_InvalidTransition(t *testing.T) {
	env := testutil.SetupTestRouter(t)
	// stream1 is "live", transition to "scheduled" is invalid
	stream1ID := "11111111-1111-1111-1111-111111111111"

	body := map[string]interface{}{
		"status": "scheduled",
	}

	resp := testutil.MakeRequest(t, http.MethodPatch, env.Server.URL+"/api/v1/streams/"+stream1ID+"/status", body)
	if resp.StatusCode != http.StatusConflict {
		t.Fatalf("expected 409, got %d", resp.StatusCode)
	}

	errResp := testutil.ParseErrorResponse(t, resp)
	if errResp.Error.Code != "INVALID_TRANSITION" {
		t.Errorf("expected error code 'INVALID_TRANSITION', got %q", errResp.Error.Code)
	}
}

func TestStreamHandler_UpdateStatus_NotFound(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	body := map[string]interface{}{"status": "live"}
	resp := testutil.MakeRequest(t, http.MethodPatch, env.Server.URL+"/api/v1/streams/00000000-0000-0000-0000-000000000000/status", body)
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", resp.StatusCode)
	}
}
