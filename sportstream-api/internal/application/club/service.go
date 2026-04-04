package club

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jpsdeveloper/sportstream-api/internal/domain"
)

type Service struct {
	repo domain.ClubRepository
}

func NewService(repo domain.ClubRepository) *Service {
	return &Service{repo: repo}
}

type CreateClubInput struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	League  string `json:"league"`
	LogoURL string `json:"logo_url"`
	Sport   string `json:"sport"`
}

type UpdateClubInput struct {
	Name     *string `json:"name,omitempty"`
	Country  *string `json:"country,omitempty"`
	League   *string `json:"league,omitempty"`
	LogoURL  *string `json:"logo_url,omitempty"`
	Sport    *string `json:"sport,omitempty"`
	IsActive *bool   `json:"is_active,omitempty"`
}

func (s *Service) List() ([]domain.Club, error) {
	return s.repo.FindAll()
}

func (s *Service) GetByID(id uuid.UUID) (*domain.Club, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Create(input CreateClubInput) (*domain.Club, error) {
	if input.Name == "" || input.Sport == "" {
		return nil, domain.ErrInvalidInput
	}

	now := time.Now().UTC()
	club := &domain.Club{
		ID:        uuid.New(),
		Name:      input.Name,
		Slug:      generateSlug(input.Name),
		Country:   input.Country,
		League:    input.League,
		LogoURL:   input.LogoURL,
		Sport:     input.Sport,
		IsActive:  true,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repo.Create(club); err != nil {
		return nil, err
	}
	return club, nil
}

func (s *Service) Update(id uuid.UUID, input UpdateClubInput) (*domain.Club, error) {
	club, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if input.Name != nil {
		club.Name = *input.Name
		club.Slug = generateSlug(*input.Name)
	}
	if input.Country != nil {
		club.Country = *input.Country
	}
	if input.League != nil {
		club.League = *input.League
	}
	if input.LogoURL != nil {
		club.LogoURL = *input.LogoURL
	}
	if input.Sport != nil {
		club.Sport = *input.Sport
	}
	if input.IsActive != nil {
		club.IsActive = *input.IsActive
	}
	club.UpdatedAt = time.Now().UTC()

	if err := s.repo.Update(club); err != nil {
		return nil, err
	}
	return club, nil
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func generateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, ".", "")
	return slug
}
