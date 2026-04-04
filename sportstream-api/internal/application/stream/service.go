package stream

import (
	"time"

	"github.com/google/uuid"
	"github.com/jpsdeveloper/sportstream-api/internal/domain"
)

type Service struct {
	repo     domain.StreamRepository
	clubRepo domain.ClubRepository
}

func NewService(repo domain.StreamRepository, clubRepo domain.ClubRepository) *Service {
	return &Service{repo: repo, clubRepo: clubRepo}
}

type CreateStreamInput struct {
	ClubID       uuid.UUID `json:"club_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Type         string    `json:"type"`
	StreamURL    string    `json:"stream_url"`
	ThumbnailURL string    `json:"thumbnail_url"`
	Duration     int       `json:"duration"`
	ScheduledAt  *string   `json:"scheduled_at,omitempty"`
	Tags         []string  `json:"tags"`
}

type UpdateStreamInput struct {
	Title        *string  `json:"title,omitempty"`
	Description  *string  `json:"description,omitempty"`
	StreamURL    *string  `json:"stream_url,omitempty"`
	ThumbnailURL *string  `json:"thumbnail_url,omitempty"`
	Duration     *int     `json:"duration,omitempty"`
	Tags         []string `json:"tags,omitempty"`
}

type UpdateStatusInput struct {
	Status string `json:"status"`
}

type StreamFilter struct {
	Status string
	Type   string
}

func (s *Service) List(filter StreamFilter) ([]domain.Stream, error) {
	streams, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	if filter.Status == "" && filter.Type == "" {
		return streams, nil
	}

	var filtered []domain.Stream
	for _, st := range streams {
		if filter.Status != "" && string(st.Status) != filter.Status {
			continue
		}
		if filter.Type != "" && string(st.Type) != filter.Type {
			continue
		}
		filtered = append(filtered, st)
	}
	return filtered, nil
}

func (s *Service) GetByID(id uuid.UUID) (*domain.Stream, error) {
	return s.repo.FindByID(id)
}

func (s *Service) GetByClubID(clubID uuid.UUID) ([]domain.Stream, error) {
	if _, err := s.clubRepo.FindByID(clubID); err != nil {
		return nil, err
	}
	return s.repo.FindByClubID(clubID)
}

func (s *Service) Create(input CreateStreamInput) (*domain.Stream, error) {
	if input.Title == "" {
		return nil, domain.ErrInvalidInput
	}
	if !domain.IsValidStreamType(input.Type) {
		return nil, domain.ErrInvalidStreamType
	}
	if _, err := s.clubRepo.FindByID(input.ClubID); err != nil {
		return nil, domain.ErrClubNotFound
	}

	now := time.Now().UTC()
	stream := &domain.Stream{
		ID:           uuid.New(),
		ClubID:       input.ClubID,
		Title:        input.Title,
		Description:  input.Description,
		Type:         domain.StreamType(input.Type),
		Status:       domain.StreamStatusScheduled,
		StreamURL:    input.StreamURL,
		ThumbnailURL: input.ThumbnailURL,
		ViewCount:    0,
		Duration:     input.Duration,
		Tags:         input.Tags,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if input.ScheduledAt != nil {
		t, err := time.Parse(time.RFC3339, *input.ScheduledAt)
		if err != nil {
			return nil, domain.ErrInvalidInput
		}
		stream.ScheduledAt = &t
	}

	if stream.Tags == nil {
		stream.Tags = []string{}
	}

	if err := s.repo.Create(stream); err != nil {
		return nil, err
	}
	return stream, nil
}

func (s *Service) Update(id uuid.UUID, input UpdateStreamInput) (*domain.Stream, error) {
	stream, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if input.Title != nil {
		stream.Title = *input.Title
	}
	if input.Description != nil {
		stream.Description = *input.Description
	}
	if input.StreamURL != nil {
		stream.StreamURL = *input.StreamURL
	}
	if input.ThumbnailURL != nil {
		stream.ThumbnailURL = *input.ThumbnailURL
	}
	if input.Duration != nil {
		stream.Duration = *input.Duration
	}
	if input.Tags != nil {
		stream.Tags = input.Tags
	}
	stream.UpdatedAt = time.Now().UTC()

	if err := s.repo.Update(stream); err != nil {
		return nil, err
	}
	return stream, nil
}

func (s *Service) UpdateStatus(id uuid.UUID, input UpdateStatusInput) (*domain.Stream, error) {
	if !domain.IsValidStreamStatus(input.Status) {
		return nil, domain.ErrInvalidStreamStatus
	}

	if err := s.repo.UpdateStatus(id, domain.StreamStatus(input.Status)); err != nil {
		return nil, err
	}

	return s.repo.FindByID(id)
}
