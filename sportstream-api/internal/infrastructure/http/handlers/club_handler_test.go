package handlers_test

import (
	"net/http"
	"testing"

	"github.com/jpsdeveloper/sportstream-api/internal/testutil"
)

func TestClubHandler_List(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/clubs", nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	result := testutil.ParseSuccessResponse(t, resp)
	data := testutil.ParseDataAsSlice(t, result)

	if len(data) != 5 {
		t.Errorf("expected 5 seeded clubs, got %d", len(data))
	}
	if result.Meta.Total != 5 {
		t.Errorf("expected meta.total 5, got %d", result.Meta.Total)
	}
}

func TestClubHandler_GetByID(t *testing.T) {
	env := testutil.SetupTestRouter(t)
	lazioID := "a1b2c3d4-e5f6-7890-abcd-ef1234567890"

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/clubs/"+lazioID, nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	result := testutil.ParseSuccessResponse(t, resp)
	data := testutil.ParseDataAsMap(t, result)

	if data["name"] != "S.S. Lazio" {
		t.Errorf("expected name 'S.S. Lazio', got %v", data["name"])
	}
}

func TestClubHandler_GetByID_InvalidID(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/clubs/not-a-uuid", nil)
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}

	errResp := testutil.ParseErrorResponse(t, resp)
	if errResp.Error.Code != "INVALID_ID" {
		t.Errorf("expected error code 'INVALID_ID', got %q", errResp.Error.Code)
	}
}

func TestClubHandler_GetByID_NotFound(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/clubs/00000000-0000-0000-0000-000000000000", nil)
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", resp.StatusCode)
	}

	errResp := testutil.ParseErrorResponse(t, resp)
	if errResp.Error.Code != "NOT_FOUND" {
		t.Errorf("expected error code 'NOT_FOUND', got %q", errResp.Error.Code)
	}
}

func TestClubHandler_Create(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	body := map[string]interface{}{
		"name":    "New Club FC",
		"country": "Germany",
		"league":  "Bundesliga",
		"sport":   "football",
	}

	resp := testutil.MakeRequest(t, http.MethodPost, env.Server.URL+"/api/v1/clubs", body)
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}

	result := testutil.ParseSuccessResponse(t, resp)
	data := testutil.ParseDataAsMap(t, result)

	if data["name"] != "New Club FC" {
		t.Errorf("expected name 'New Club FC', got %v", data["name"])
	}
	if data["slug"] != "new-club-fc" {
		t.Errorf("expected slug 'new-club-fc', got %v", data["slug"])
	}
	if data["is_active"] != true {
		t.Errorf("expected is_active true, got %v", data["is_active"])
	}
}

func TestClubHandler_Create_InvalidBody(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	// Missing required fields
	body := map[string]interface{}{
		"country": "Germany",
	}

	resp := testutil.MakeRequest(t, http.MethodPost, env.Server.URL+"/api/v1/clubs", body)
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}

func TestClubHandler_Update(t *testing.T) {
	env := testutil.SetupTestRouter(t)
	lazioID := "a1b2c3d4-e5f6-7890-abcd-ef1234567890"

	body := map[string]interface{}{
		"name": "S.S. Lazio Updated",
	}

	resp := testutil.MakeRequest(t, http.MethodPut, env.Server.URL+"/api/v1/clubs/"+lazioID, body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	result := testutil.ParseSuccessResponse(t, resp)
	data := testutil.ParseDataAsMap(t, result)

	if data["name"] != "S.S. Lazio Updated" {
		t.Errorf("expected name 'S.S. Lazio Updated', got %v", data["name"])
	}
}

func TestClubHandler_Update_NotFound(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	body := map[string]interface{}{"name": "Whatever"}
	resp := testutil.MakeRequest(t, http.MethodPut, env.Server.URL+"/api/v1/clubs/00000000-0000-0000-0000-000000000000", body)
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", resp.StatusCode)
	}
}

func TestClubHandler_Delete(t *testing.T) {
	env := testutil.SetupTestRouter(t)
	lazioID := "a1b2c3d4-e5f6-7890-abcd-ef1234567890"

	resp := testutil.MakeRequest(t, http.MethodDelete, env.Server.URL+"/api/v1/clubs/"+lazioID, nil)
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", resp.StatusCode)
	}

	// Verify it's gone
	resp2 := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/clubs/"+lazioID, nil)
	if resp2.StatusCode != http.StatusNotFound {
		resp2.Body.Close()
		t.Fatalf("expected 404 after delete, got %d", resp2.StatusCode)
	}
	resp2.Body.Close()
}

func TestClubHandler_Delete_NotFound(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	resp := testutil.MakeRequest(t, http.MethodDelete, env.Server.URL+"/api/v1/clubs/00000000-0000-0000-0000-000000000000", nil)
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", resp.StatusCode)
	}
}
