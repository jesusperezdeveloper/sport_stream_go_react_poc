package club

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/jpsdeveloper/sportstream-api/internal/domain"
	"github.com/jpsdeveloper/sportstream-api/internal/infrastructure/persistence/memory"
)

func setupClubService() *Service {
	repo := memory.NewClubRepository()
	return NewService(repo)
}

func TestClubService_Create_Valid(t *testing.T) {
	svc := setupClubService()

	club, err := svc.Create(CreateClubInput{
		Name:    "Test FC",
		Country: "Spain",
		League:  "La Liga",
		Sport:   "football",
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if club.Name != "Test FC" {
		t.Errorf("expected name 'Test FC', got %q", club.Name)
	}
	if club.Slug != "test-fc" {
		t.Errorf("expected slug 'test-fc', got %q", club.Slug)
	}
	if !club.IsActive {
		t.Error("expected club to be active")
	}
	if club.ID == uuid.Nil {
		t.Error("expected non-nil UUID")
	}
}

func TestClubService_Create_MissingName(t *testing.T) {
	svc := setupClubService()

	_, err := svc.Create(CreateClubInput{
		Sport: "football",
	})
	if !errors.Is(err, domain.ErrInvalidInput) {
		t.Errorf("expected ErrInvalidInput, got %v", err)
	}
}

func TestClubService_Create_MissingSport(t *testing.T) {
	svc := setupClubService()

	_, err := svc.Create(CreateClubInput{
		Name: "Test FC",
	})
	if !errors.Is(err, domain.ErrInvalidInput) {
		t.Errorf("expected ErrInvalidInput, got %v", err)
	}
}

func TestClubService_List(t *testing.T) {
	svc := setupClubService()

	// Create two clubs
	_, _ = svc.Create(CreateClubInput{Name: "Club A", Sport: "football"})
	_, _ = svc.Create(CreateClubInput{Name: "Club B", Sport: "tennis"})

	clubs, err := svc.List()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(clubs) != 2 {
		t.Errorf("expected 2 clubs, got %d", len(clubs))
	}
}

func TestClubService_GetByID(t *testing.T) {
	svc := setupClubService()

	created, _ := svc.Create(CreateClubInput{Name: "Club A", Sport: "football"})

	found, err := svc.GetByID(created.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if found.Name != "Club A" {
		t.Errorf("expected name 'Club A', got %q", found.Name)
	}
}

func TestClubService_GetByID_NotFound(t *testing.T) {
	svc := setupClubService()

	_, err := svc.GetByID(uuid.New())
	if !errors.Is(err, domain.ErrClubNotFound) {
		t.Errorf("expected ErrClubNotFound, got %v", err)
	}
}

func TestClubService_Update(t *testing.T) {
	svc := setupClubService()

	created, _ := svc.Create(CreateClubInput{Name: "Club A", Sport: "football", Country: "Spain"})

	newName := "Club A Updated"
	newCountry := "Italy"
	updated, err := svc.Update(created.ID, UpdateClubInput{
		Name:    &newName,
		Country: &newCountry,
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if updated.Name != "Club A Updated" {
		t.Errorf("expected name 'Club A Updated', got %q", updated.Name)
	}
	if updated.Country != "Italy" {
		t.Errorf("expected country 'Italy', got %q", updated.Country)
	}
	if updated.Slug != "club-a-updated" {
		t.Errorf("expected slug 'club-a-updated', got %q", updated.Slug)
	}
}

func TestClubService_Update_NotFound(t *testing.T) {
	svc := setupClubService()

	newName := "Whatever"
	_, err := svc.Update(uuid.New(), UpdateClubInput{Name: &newName})
	if !errors.Is(err, domain.ErrClubNotFound) {
		t.Errorf("expected ErrClubNotFound, got %v", err)
	}
}

func TestClubService_Delete(t *testing.T) {
	svc := setupClubService()

	created, _ := svc.Create(CreateClubInput{Name: "Club A", Sport: "football"})

	err := svc.Delete(created.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = svc.GetByID(created.ID)
	if !errors.Is(err, domain.ErrClubNotFound) {
		t.Errorf("expected ErrClubNotFound after delete, got %v", err)
	}
}

func TestClubService_Delete_NotFound(t *testing.T) {
	svc := setupClubService()

	err := svc.Delete(uuid.New())
	if !errors.Is(err, domain.ErrClubNotFound) {
		t.Errorf("expected ErrClubNotFound, got %v", err)
	}
}

func TestClubService_Update_IsActive(t *testing.T) {
	svc := setupClubService()

	created, _ := svc.Create(CreateClubInput{Name: "Club A", Sport: "football"})
	if !created.IsActive {
		t.Fatal("expected club to be active initially")
	}

	inactive := false
	updated, err := svc.Update(created.ID, UpdateClubInput{IsActive: &inactive})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if updated.IsActive {
		t.Error("expected club to be inactive after update")
	}
}
