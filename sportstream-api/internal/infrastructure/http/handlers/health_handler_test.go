package handlers_test

import (
	"net/http"
	"testing"

	"github.com/jpsdeveloper/sportstream-api/internal/testutil"
)

func TestHealthEndpoint(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/health", nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	result := testutil.ParseSuccessResponse(t, resp)
	data := testutil.ParseDataAsMap(t, result)

	if data["status"] != "ok" {
		t.Errorf("expected status 'ok', got %v", data["status"])
	}
	if data["version"] != "test-v1" {
		t.Errorf("expected version 'test-v1', got %v", data["version"])
	}
}
