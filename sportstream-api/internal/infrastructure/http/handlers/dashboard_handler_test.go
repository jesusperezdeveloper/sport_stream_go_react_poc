package handlers_test

import (
	"net/http"
	"testing"

	"github.com/jpsdeveloper/sportstream-api/internal/testutil"
)

func TestDashboardHandler_Summary(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/dashboard/summary", nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	result := testutil.ParseSuccessResponse(t, resp)
	data := testutil.ParseDataAsMap(t, result)

	// Verify aggregation fields exist and have correct types
	totalClubs, ok := data["total_clubs"].(float64)
	if !ok {
		t.Fatal("expected total_clubs to be a number")
	}
	if totalClubs != 5 {
		t.Errorf("expected 5 total clubs, got %v", totalClubs)
	}

	totalStreams, ok := data["total_streams"].(float64)
	if !ok {
		t.Fatal("expected total_streams to be a number")
	}
	if totalStreams != 10 {
		t.Errorf("expected 10 total streams, got %v", totalStreams)
	}

	liveStreams, ok := data["live_streams"].(float64)
	if !ok {
		t.Fatal("expected live_streams to be a number")
	}
	if liveStreams != 2 {
		t.Errorf("expected 2 live streams, got %v", liveStreams)
	}

	upcomingEvents, ok := data["upcoming_events"].(float64)
	if !ok {
		t.Fatal("expected upcoming_events to be a number")
	}
	if upcomingEvents != 4 {
		t.Errorf("expected 4 upcoming events, got %v", upcomingEvents)
	}

	totalViews, ok := data["total_views"].(float64)
	if !ok {
		t.Fatal("expected total_views to be a number")
	}
	if totalViews <= 0 {
		t.Error("expected positive total views")
	}

	// Verify streams_by_type is a map
	streamsByType, ok := data["streams_by_type"].(map[string]interface{})
	if !ok {
		t.Fatal("expected streams_by_type to be a map")
	}
	if len(streamsByType) == 0 {
		t.Error("expected streams_by_type to be non-empty")
	}

	// Verify streams_by_status is a map
	streamsByStatus, ok := data["streams_by_status"].(map[string]interface{})
	if !ok {
		t.Fatal("expected streams_by_status to be a map")
	}
	if len(streamsByStatus) == 0 {
		t.Error("expected streams_by_status to be non-empty")
	}

	// Verify top_clubs_by_views is a slice
	topClubs, ok := data["top_clubs_by_views"].([]interface{})
	if !ok {
		t.Fatal("expected top_clubs_by_views to be a slice")
	}
	if len(topClubs) == 0 {
		t.Error("expected top_clubs_by_views to be non-empty")
	}
}

func TestDashboardHandler_Summary_ResponseFormat(t *testing.T) {
	env := testutil.SetupTestRouter(t)

	resp := testutil.MakeRequest(t, http.MethodGet, env.Server.URL+"/api/v1/dashboard/summary", nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	result := testutil.ParseSuccessResponse(t, resp)

	if result.Meta.Total != 1 {
		t.Errorf("expected meta.total 1, got %d", result.Meta.Total)
	}
	if result.Meta.Timestamp == "" {
		t.Error("expected meta.timestamp to be set")
	}
}
