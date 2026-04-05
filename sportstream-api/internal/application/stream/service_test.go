package stream

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/jpsdeveloper/sportstream-api/internal/domain"
	"github.com/jpsdeveloper/sportstream-api/internal/infrastructure/persistence/memory"
)

func setupStreamService() (*Service, uuid.UUID) {
	clubRepo := memory.NewClubRepository()
	streamRepo := memory.NewStreamRepository()

	// Create a club for stream references
	clubID := uuid.New()
	_ = clubRepo.Create(&domain.Club{
		ID:       clubID,
		Name:     "Test Club",
		Sport:    "football",
		IsActive: true,
	})

	svc := NewService(streamRepo, clubRepo)
	return svc, clubID
}

func TestStreamService_Create_Valid(t *testing.T) {
	svc, clubID := setupStreamService()

	stream, err := svc.Create(CreateStreamInput{
		ClubID:  clubID,
		Title:   "Test Match",
		Type:    "live",
		Tags:    []string{"football"},
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if stream.Title != "Test Match" {
		t.Errorf("expected title 'Test Match', got %q", stream.Title)
	}
	if stream.Status != domain.StreamStatusScheduled {
		t.Errorf("expected status 'scheduled', got %q", stream.Status)
	}
	if stream.ID == uuid.Nil {
		t.Error("expected non-nil UUID")
	}
}

func TestStreamService_Create_MissingTitle(t *testing.T) {
	svc, clubID := setupStreamService()

	_, err := svc.Create(CreateStreamInput{
		ClubID: clubID,
		Type:   "live",
	})
	if !errors.Is(err, domain.ErrInvalidInput) {
		t.Errorf("expected ErrInvalidInput, got %v", err)
	}
}

func TestStreamService_Create_InvalidType(t *testing.T) {
	svc, clubID := setupStreamService()

	_, err := svc.Create(CreateStreamInput{
		ClubID: clubID,
		Title:  "Test",
		Type:   "podcast",
	})
	if !errors.Is(err, domain.ErrInvalidStreamType) {
		t.Errorf("expected ErrInvalidStreamType, got %v", err)
	}
}

func TestStreamService_Create_ClubNotFound(t *testing.T) {
	svc, _ := setupStreamService()

	_, err := svc.Create(CreateStreamInput{
		ClubID: uuid.New(),
		Title:  "Test",
		Type:   "live",
	})
	if !errors.Is(err, domain.ErrClubNotFound) {
		t.Errorf("expected ErrClubNotFound, got %v", err)
	}
}

func TestStreamService_Create_WithScheduledAt(t *testing.T) {
	svc, clubID := setupStreamService()

	stream, err := svc.Create(CreateStreamInput{
		ClubID: clubID,
		Title:  "Scheduled Match",
		Type:   "live",
		ScheduledAt: strPtr("2026-06-01T15:00:00Z"),
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if stream.ScheduledAt == nil {
		t.Error("expected ScheduledAt to be set")
	}
}

func TestStreamService_Create_InvalidScheduledAt(t *testing.T) {
	svc, clubID := setupStreamService()

	_, err := svc.Create(CreateStreamInput{
		ClubID:      clubID,
		Title:       "Bad Date",
		Type:        "live",
		ScheduledAt: strPtr("not-a-date"),
	})
	if !errors.Is(err, domain.ErrInvalidInput) {
		t.Errorf("expected ErrInvalidInput, got %v", err)
	}
}

func TestStreamService_Create_NilTagsBecomesEmptySlice(t *testing.T) {
	svc, clubID := setupStreamService()

	stream, err := svc.Create(CreateStreamInput{
		ClubID: clubID,
		Title:  "No Tags",
		Type:   "vod",
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if stream.Tags == nil {
		t.Error("expected Tags to be non-nil empty slice")
	}
	if len(stream.Tags) != 0 {
		t.Errorf("expected 0 tags, got %d", len(stream.Tags))
	}
}

func TestStreamService_List(t *testing.T) {
	svc, clubID := setupStreamService()

	_, _ = svc.Create(CreateStreamInput{ClubID: clubID, Title: "Stream A", Type: "live"})
	_, _ = svc.Create(CreateStreamInput{ClubID: clubID, Title: "Stream B", Type: "vod"})

	streams, err := svc.List(StreamFilter{})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(streams) != 2 {
		t.Errorf("expected 2 streams, got %d", len(streams))
	}
}

func TestStreamService_List_FilterByStatus(t *testing.T) {
	svc, clubID := setupStreamService()

	_, _ = svc.Create(CreateStreamInput{ClubID: clubID, Title: "Stream A", Type: "live"})
	_, _ = svc.Create(CreateStreamInput{ClubID: clubID, Title: "Stream B", Type: "vod"})

	// All new streams are "scheduled"
	streams, err := svc.List(StreamFilter{Status: "scheduled"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(streams) != 2 {
		t.Errorf("expected 2 scheduled streams, got %d", len(streams))
	}

	streams, err = svc.List(StreamFilter{Status: "live"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(streams) != 0 {
		t.Errorf("expected 0 live streams, got %d", len(streams))
	}
}

func TestStreamService_List_FilterByType(t *testing.T) {
	svc, clubID := setupStreamService()

	_, _ = svc.Create(CreateStreamInput{ClubID: clubID, Title: "Stream A", Type: "live"})
	_, _ = svc.Create(CreateStreamInput{ClubID: clubID, Title: "Stream B", Type: "vod"})

	streams, err := svc.List(StreamFilter{Type: "vod"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(streams) != 1 {
		t.Errorf("expected 1 vod stream, got %d", len(streams))
	}
}

func TestStreamService_GetByID(t *testing.T) {
	svc, clubID := setupStreamService()

	created, _ := svc.Create(CreateStreamInput{ClubID: clubID, Title: "Stream A", Type: "live"})

	found, err := svc.GetByID(created.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if found.Title != "Stream A" {
		t.Errorf("expected title 'Stream A', got %q", found.Title)
	}
}

func TestStreamService_GetByID_NotFound(t *testing.T) {
	svc, _ := setupStreamService()

	_, err := svc.GetByID(uuid.New())
	if !errors.Is(err, domain.ErrStreamNotFound) {
		t.Errorf("expected ErrStreamNotFound, got %v", err)
	}
}

func TestStreamService_GetByClubID(t *testing.T) {
	svc, clubID := setupStreamService()

	_, _ = svc.Create(CreateStreamInput{ClubID: clubID, Title: "Stream A", Type: "live"})
	_, _ = svc.Create(CreateStreamInput{ClubID: clubID, Title: "Stream B", Type: "vod"})

	streams, err := svc.GetByClubID(clubID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(streams) != 2 {
		t.Errorf("expected 2 streams for club, got %d", len(streams))
	}
}

func TestStreamService_GetByClubID_ClubNotFound(t *testing.T) {
	svc, _ := setupStreamService()

	_, err := svc.GetByClubID(uuid.New())
	if !errors.Is(err, domain.ErrClubNotFound) {
		t.Errorf("expected ErrClubNotFound, got %v", err)
	}
}

func TestStreamService_Update(t *testing.T) {
	svc, clubID := setupStreamService()

	created, _ := svc.Create(CreateStreamInput{ClubID: clubID, Title: "Stream A", Type: "live"})

	newTitle := "Stream A Updated"
	newDuration := 3600
	updated, err := svc.Update(created.ID, UpdateStreamInput{
		Title:    &newTitle,
		Duration: &newDuration,
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if updated.Title != "Stream A Updated" {
		t.Errorf("expected title 'Stream A Updated', got %q", updated.Title)
	}
	if updated.Duration != 3600 {
		t.Errorf("expected duration 3600, got %d", updated.Duration)
	}
}

func TestStreamService_Update_NotFound(t *testing.T) {
	svc, _ := setupStreamService()

	newTitle := "Whatever"
	_, err := svc.Update(uuid.New(), UpdateStreamInput{Title: &newTitle})
	if !errors.Is(err, domain.ErrStreamNotFound) {
		t.Errorf("expected ErrStreamNotFound, got %v", err)
	}
}

func TestStreamService_UpdateStatus_ValidTransition(t *testing.T) {
	svc, clubID := setupStreamService()

	created, _ := svc.Create(CreateStreamInput{ClubID: clubID, Title: "Stream A", Type: "live"})

	// scheduled -> live
	updated, err := svc.UpdateStatus(created.ID, UpdateStatusInput{Status: "live"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if updated.Status != domain.StreamStatusLive {
		t.Errorf("expected status 'live', got %q", updated.Status)
	}
}

func TestStreamService_UpdateStatus_InvalidStatus(t *testing.T) {
	svc, clubID := setupStreamService()

	created, _ := svc.Create(CreateStreamInput{ClubID: clubID, Title: "Stream A", Type: "live"})

	_, err := svc.UpdateStatus(created.ID, UpdateStatusInput{Status: "bogus"})
	if !errors.Is(err, domain.ErrInvalidStreamStatus) {
		t.Errorf("expected ErrInvalidStreamStatus, got %v", err)
	}
}

func TestStreamService_UpdateStatus_InvalidTransition(t *testing.T) {
	svc, clubID := setupStreamService()

	created, _ := svc.Create(CreateStreamInput{ClubID: clubID, Title: "Stream A", Type: "live"})

	// scheduled -> ended is invalid (must go through live first)
	_, err := svc.UpdateStatus(created.ID, UpdateStatusInput{Status: "ended"})
	if !errors.Is(err, domain.ErrInvalidTransition) {
		t.Errorf("expected ErrInvalidTransition, got %v", err)
	}
}

func TestStreamService_UpdateStatus_NotFound(t *testing.T) {
	svc, _ := setupStreamService()

	_, err := svc.UpdateStatus(uuid.New(), UpdateStatusInput{Status: "live"})
	if !errors.Is(err, domain.ErrStreamNotFound) {
		t.Errorf("expected ErrStreamNotFound, got %v", err)
	}
}

func strPtr(s string) *string {
	return &s
}
