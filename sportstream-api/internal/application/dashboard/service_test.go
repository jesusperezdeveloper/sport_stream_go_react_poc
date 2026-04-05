package dashboard

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jpsdeveloper/sportstream-api/internal/domain"
	"github.com/jpsdeveloper/sportstream-api/internal/infrastructure/persistence/memory"
)

func setupDashboardService() (*Service, *memory.ClubRepository, *memory.StreamRepository, *memory.EventRepository) {
	clubRepo := memory.NewClubRepository()
	streamRepo := memory.NewStreamRepository()
	eventRepo := memory.NewEventRepository()
	svc := NewService(clubRepo, streamRepo, eventRepo)
	return svc, clubRepo, streamRepo, eventRepo
}

func TestDashboardService_GetSummary_Empty(t *testing.T) {
	svc, _, _, _ := setupDashboardService()

	summary, err := svc.GetSummary()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if summary.TotalClubs != 0 {
		t.Errorf("expected 0 clubs, got %d", summary.TotalClubs)
	}
	if summary.TotalStreams != 0 {
		t.Errorf("expected 0 streams, got %d", summary.TotalStreams)
	}
	if summary.LiveStreams != 0 {
		t.Errorf("expected 0 live streams, got %d", summary.LiveStreams)
	}
	if summary.UpcomingEvents != 0 {
		t.Errorf("expected 0 upcoming events, got %d", summary.UpcomingEvents)
	}
	if summary.TotalViews != 0 {
		t.Errorf("expected 0 total views, got %d", summary.TotalViews)
	}
}

func TestDashboardService_GetSummary_WithData(t *testing.T) {
	svc, clubRepo, streamRepo, eventRepo := setupDashboardService()

	now := time.Now().UTC()
	clubA := uuid.New()
	clubB := uuid.New()

	_ = clubRepo.Create(&domain.Club{ID: clubA, Name: "Club A", Sport: "football", CreatedAt: now, UpdatedAt: now})
	_ = clubRepo.Create(&domain.Club{ID: clubB, Name: "Club B", Sport: "tennis", CreatedAt: now, UpdatedAt: now})

	_ = streamRepo.Create(&domain.Stream{
		ID: uuid.New(), ClubID: clubA, Title: "Stream 1",
		Type: domain.StreamTypeLive, Status: domain.StreamStatusLive,
		ViewCount: 1000, Tags: []string{}, CreatedAt: now, UpdatedAt: now,
	})
	_ = streamRepo.Create(&domain.Stream{
		ID: uuid.New(), ClubID: clubA, Title: "Stream 2",
		Type: domain.StreamTypeVOD, Status: domain.StreamStatusArchived,
		ViewCount: 500, Tags: []string{}, CreatedAt: now, UpdatedAt: now,
	})
	_ = streamRepo.Create(&domain.Stream{
		ID: uuid.New(), ClubID: clubB, Title: "Stream 3",
		Type: domain.StreamTypeLive, Status: domain.StreamStatusScheduled,
		ViewCount: 0, Tags: []string{}, CreatedAt: now, UpdatedAt: now,
	})

	_ = eventRepo.Create(&domain.Event{
		ID: uuid.New(), ClubID: clubA, Title: "Event 1",
		Sport: "football", Status: domain.EventStatusUpcoming,
		StartTime: now.Add(24 * time.Hour), CreatedAt: now, UpdatedAt: now,
	})
	_ = eventRepo.Create(&domain.Event{
		ID: uuid.New(), ClubID: clubB, Title: "Event 2",
		Sport: "tennis", Status: domain.EventStatusLive,
		StartTime: now, CreatedAt: now, UpdatedAt: now,
	})

	summary, err := svc.GetSummary()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if summary.TotalClubs != 2 {
		t.Errorf("expected 2 clubs, got %d", summary.TotalClubs)
	}
	if summary.TotalStreams != 3 {
		t.Errorf("expected 3 streams, got %d", summary.TotalStreams)
	}
	if summary.LiveStreams != 1 {
		t.Errorf("expected 1 live stream, got %d", summary.LiveStreams)
	}
	if summary.UpcomingEvents != 1 {
		t.Errorf("expected 1 upcoming event, got %d", summary.UpcomingEvents)
	}
	if summary.TotalViews != 1500 {
		t.Errorf("expected 1500 total views, got %d", summary.TotalViews)
	}

	// Check streams by type
	if summary.StreamsByType["live"] != 2 {
		t.Errorf("expected 2 live-type streams, got %d", summary.StreamsByType["live"])
	}
	if summary.StreamsByType["vod"] != 1 {
		t.Errorf("expected 1 vod-type stream, got %d", summary.StreamsByType["vod"])
	}

	// Check streams by status
	if summary.StreamsByStatus["live"] != 1 {
		t.Errorf("expected 1 live-status stream, got %d", summary.StreamsByStatus["live"])
	}
	if summary.StreamsByStatus["archived"] != 1 {
		t.Errorf("expected 1 archived-status stream, got %d", summary.StreamsByStatus["archived"])
	}
	if summary.StreamsByStatus["scheduled"] != 1 {
		t.Errorf("expected 1 scheduled-status stream, got %d", summary.StreamsByStatus["scheduled"])
	}

	// Check top clubs by views
	if len(summary.TopClubsByViews) != 2 {
		t.Fatalf("expected 2 top clubs, got %d", len(summary.TopClubsByViews))
	}
	if summary.TopClubsByViews[0].ClubName != "Club A" {
		t.Errorf("expected top club to be 'Club A', got %q", summary.TopClubsByViews[0].ClubName)
	}
	if summary.TopClubsByViews[0].TotalViews != 1500 {
		t.Errorf("expected top club views 1500, got %d", summary.TopClubsByViews[0].TotalViews)
	}
}

func TestDashboardService_GetSummary_WithSeedData(t *testing.T) {
	clubRepo := memory.NewClubRepository()
	streamRepo := memory.NewStreamRepository()
	eventRepo := memory.NewEventRepository()
	memory.SeedData(clubRepo, streamRepo, eventRepo)

	svc := NewService(clubRepo, streamRepo, eventRepo)

	summary, err := svc.GetSummary()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if summary.TotalClubs != 5 {
		t.Errorf("expected 5 clubs, got %d", summary.TotalClubs)
	}
	if summary.TotalStreams != 10 {
		t.Errorf("expected 10 streams, got %d", summary.TotalStreams)
	}
	if summary.LiveStreams != 2 {
		t.Errorf("expected 2 live streams, got %d", summary.LiveStreams)
	}
	// Seed data has 4 upcoming events with future start times
	if summary.UpcomingEvents != 4 {
		t.Errorf("expected 4 upcoming events, got %d", summary.UpcomingEvents)
	}
	if summary.TotalViews <= 0 {
		t.Error("expected positive total views")
	}
}
