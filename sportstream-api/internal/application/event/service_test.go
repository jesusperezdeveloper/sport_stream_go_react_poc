package event

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jpsdeveloper/sportstream-api/internal/domain"
	"github.com/jpsdeveloper/sportstream-api/internal/infrastructure/persistence/memory"
)

func setupEventService() (*Service, uuid.UUID) {
	clubRepo := memory.NewClubRepository()
	eventRepo := memory.NewEventRepository()

	clubID := uuid.New()
	_ = clubRepo.Create(&domain.Club{
		ID:       clubID,
		Name:     "Test Club",
		Sport:    "football",
		IsActive: true,
	})

	svc := NewService(eventRepo, clubRepo)
	return svc, clubID
}

func TestEventService_Create_Valid(t *testing.T) {
	svc, clubID := setupEventService()

	event, err := svc.Create(CreateEventInput{
		ClubID:    clubID,
		Title:     "Test Match",
		Sport:     "football",
		Venue:     "Test Stadium",
		StartTime: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if event.Title != "Test Match" {
		t.Errorf("expected title 'Test Match', got %q", event.Title)
	}
	if event.Status != domain.EventStatusUpcoming {
		t.Errorf("expected status 'upcoming', got %q", event.Status)
	}
	if event.ID == uuid.Nil {
		t.Error("expected non-nil UUID")
	}
}

func TestEventService_Create_MissingTitle(t *testing.T) {
	svc, clubID := setupEventService()

	_, err := svc.Create(CreateEventInput{
		ClubID:    clubID,
		Sport:     "football",
		StartTime: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	})
	if !errors.Is(err, domain.ErrInvalidInput) {
		t.Errorf("expected ErrInvalidInput, got %v", err)
	}
}

func TestEventService_Create_MissingSport(t *testing.T) {
	svc, clubID := setupEventService()

	_, err := svc.Create(CreateEventInput{
		ClubID:    clubID,
		Title:     "Test Match",
		StartTime: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	})
	if !errors.Is(err, domain.ErrInvalidInput) {
		t.Errorf("expected ErrInvalidInput, got %v", err)
	}
}

func TestEventService_Create_ClubNotFound(t *testing.T) {
	svc, _ := setupEventService()

	_, err := svc.Create(CreateEventInput{
		ClubID:    uuid.New(),
		Title:     "Test",
		Sport:     "football",
		StartTime: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	})
	if !errors.Is(err, domain.ErrClubNotFound) {
		t.Errorf("expected ErrClubNotFound, got %v", err)
	}
}

func TestEventService_Create_InvalidStartTime(t *testing.T) {
	svc, clubID := setupEventService()

	_, err := svc.Create(CreateEventInput{
		ClubID:    clubID,
		Title:     "Test",
		Sport:     "football",
		StartTime: "not-a-date",
	})
	if !errors.Is(err, domain.ErrInvalidInput) {
		t.Errorf("expected ErrInvalidInput, got %v", err)
	}
}

func TestEventService_Create_WithEndTime(t *testing.T) {
	svc, clubID := setupEventService()

	endTime := time.Now().Add(26 * time.Hour).Format(time.RFC3339)
	event, err := svc.Create(CreateEventInput{
		ClubID:    clubID,
		Title:     "Test",
		Sport:     "football",
		StartTime: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
		EndTime:   &endTime,
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if event.EndTime == nil {
		t.Error("expected EndTime to be set")
	}
}

func TestEventService_Create_InvalidEndTime(t *testing.T) {
	svc, clubID := setupEventService()

	badEndTime := "not-a-date"
	_, err := svc.Create(CreateEventInput{
		ClubID:    clubID,
		Title:     "Test",
		Sport:     "football",
		StartTime: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
		EndTime:   &badEndTime,
	})
	if !errors.Is(err, domain.ErrInvalidInput) {
		t.Errorf("expected ErrInvalidInput, got %v", err)
	}
}

func TestEventService_List(t *testing.T) {
	svc, clubID := setupEventService()

	_, _ = svc.Create(CreateEventInput{
		ClubID: clubID, Title: "Event A", Sport: "football",
		StartTime: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	})
	_, _ = svc.Create(CreateEventInput{
		ClubID: clubID, Title: "Event B", Sport: "tennis",
		StartTime: time.Now().Add(48 * time.Hour).Format(time.RFC3339),
	})

	events, err := svc.List(EventFilter{})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(events) != 2 {
		t.Errorf("expected 2 events, got %d", len(events))
	}
}

func TestEventService_List_FilterByStatus(t *testing.T) {
	svc, clubID := setupEventService()

	_, _ = svc.Create(CreateEventInput{
		ClubID: clubID, Title: "Event A", Sport: "football",
		StartTime: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	})

	events, err := svc.List(EventFilter{Status: "upcoming"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(events) != 1 {
		t.Errorf("expected 1 upcoming event, got %d", len(events))
	}

	events, err = svc.List(EventFilter{Status: "live"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(events) != 0 {
		t.Errorf("expected 0 live events, got %d", len(events))
	}
}

func TestEventService_List_FilterBySport(t *testing.T) {
	svc, clubID := setupEventService()

	_, _ = svc.Create(CreateEventInput{
		ClubID: clubID, Title: "Event A", Sport: "football",
		StartTime: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	})
	_, _ = svc.Create(CreateEventInput{
		ClubID: clubID, Title: "Event B", Sport: "tennis",
		StartTime: time.Now().Add(48 * time.Hour).Format(time.RFC3339),
	})

	events, err := svc.List(EventFilter{Sport: "tennis"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(events) != 1 {
		t.Errorf("expected 1 tennis event, got %d", len(events))
	}
}

func TestEventService_GetByID(t *testing.T) {
	svc, clubID := setupEventService()

	created, _ := svc.Create(CreateEventInput{
		ClubID: clubID, Title: "Event A", Sport: "football",
		StartTime: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	})

	found, err := svc.GetByID(created.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if found.Title != "Event A" {
		t.Errorf("expected title 'Event A', got %q", found.Title)
	}
}

func TestEventService_GetByID_NotFound(t *testing.T) {
	svc, _ := setupEventService()

	_, err := svc.GetByID(uuid.New())
	if !errors.Is(err, domain.ErrEventNotFound) {
		t.Errorf("expected ErrEventNotFound, got %v", err)
	}
}

func TestEventService_GetUpcoming(t *testing.T) {
	svc, clubID := setupEventService()

	// Create an event in the future
	_, _ = svc.Create(CreateEventInput{
		ClubID: clubID, Title: "Future Event", Sport: "football",
		StartTime: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	})

	events, err := svc.GetUpcoming()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(events) != 1 {
		t.Errorf("expected 1 upcoming event, got %d", len(events))
	}
}

func TestEventService_Update(t *testing.T) {
	svc, clubID := setupEventService()

	created, _ := svc.Create(CreateEventInput{
		ClubID: clubID, Title: "Event A", Sport: "football", Venue: "Stadium A",
		StartTime: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	})

	updated, err := svc.Update(created.ID, CreateEventInput{
		ClubID: clubID, Title: "Event A Updated", Sport: "football", Venue: "Stadium B",
		StartTime: time.Now().Add(48 * time.Hour).Format(time.RFC3339),
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if updated.Title != "Event A Updated" {
		t.Errorf("expected title 'Event A Updated', got %q", updated.Title)
	}
	if updated.Venue != "Stadium B" {
		t.Errorf("expected venue 'Stadium B', got %q", updated.Venue)
	}
}

func TestEventService_Update_NotFound(t *testing.T) {
	svc, clubID := setupEventService()

	_, err := svc.Update(uuid.New(), CreateEventInput{
		ClubID: clubID, Title: "Event A", Sport: "football",
		StartTime: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	})
	if !errors.Is(err, domain.ErrEventNotFound) {
		t.Errorf("expected ErrEventNotFound, got %v", err)
	}
}

func TestEventService_Update_InvalidStartTime(t *testing.T) {
	svc, clubID := setupEventService()

	created, _ := svc.Create(CreateEventInput{
		ClubID: clubID, Title: "Event A", Sport: "football",
		StartTime: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	})

	_, err := svc.Update(created.ID, CreateEventInput{
		ClubID: clubID, Title: "Event A", Sport: "football",
		StartTime: "not-a-date",
	})
	if !errors.Is(err, domain.ErrInvalidInput) {
		t.Errorf("expected ErrInvalidInput, got %v", err)
	}
}
