package event

import (
	"time"

	"github.com/google/uuid"
	"github.com/jpsdeveloper/sportstream-api/internal/domain"
)

type Service struct {
	repo     domain.EventRepository
	clubRepo domain.ClubRepository
}

func NewService(repo domain.EventRepository, clubRepo domain.ClubRepository) *Service {
	return &Service{repo: repo, clubRepo: clubRepo}
}

type CreateEventInput struct {
	ClubID      uuid.UUID  `json:"club_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Venue       string     `json:"venue"`
	Sport       string     `json:"sport"`
	StartTime   string     `json:"start_time"`
	EndTime     *string    `json:"end_time,omitempty"`
	StreamID    *uuid.UUID `json:"stream_id,omitempty"`
}

type EventFilter struct {
	Status string
	Sport  string
}

func (s *Service) List(filter EventFilter) ([]domain.Event, error) {
	events, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	if filter.Status == "" && filter.Sport == "" {
		return events, nil
	}

	var filtered []domain.Event
	for _, e := range events {
		if filter.Status != "" && string(e.Status) != filter.Status {
			continue
		}
		if filter.Sport != "" && e.Sport != filter.Sport {
			continue
		}
		filtered = append(filtered, e)
	}
	return filtered, nil
}

func (s *Service) GetByID(id uuid.UUID) (*domain.Event, error) {
	return s.repo.FindByID(id)
}

func (s *Service) GetUpcoming() ([]domain.Event, error) {
	return s.repo.FindUpcoming()
}

func (s *Service) Create(input CreateEventInput) (*domain.Event, error) {
	if input.Title == "" || input.Sport == "" {
		return nil, domain.ErrInvalidInput
	}
	if _, err := s.clubRepo.FindByID(input.ClubID); err != nil {
		return nil, domain.ErrClubNotFound
	}

	startTime, err := time.Parse(time.RFC3339, input.StartTime)
	if err != nil {
		return nil, domain.ErrInvalidInput
	}

	now := time.Now().UTC()
	event := &domain.Event{
		ID:          uuid.New(),
		ClubID:      input.ClubID,
		Title:       input.Title,
		Description: input.Description,
		Venue:       input.Venue,
		Sport:       input.Sport,
		StartTime:   startTime,
		Status:      domain.EventStatusUpcoming,
		StreamID:    input.StreamID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if input.EndTime != nil {
		t, err := time.Parse(time.RFC3339, *input.EndTime)
		if err != nil {
			return nil, domain.ErrInvalidInput
		}
		event.EndTime = &t
	}

	if err := s.repo.Create(event); err != nil {
		return nil, err
	}
	return event, nil
}

func (s *Service) Update(id uuid.UUID, input CreateEventInput) (*domain.Event, error) {
	event, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	startTime, err := time.Parse(time.RFC3339, input.StartTime)
	if err != nil {
		return nil, domain.ErrInvalidInput
	}

	event.Title = input.Title
	event.Description = input.Description
	event.Venue = input.Venue
	event.Sport = input.Sport
	event.StartTime = startTime
	event.StreamID = input.StreamID
	event.UpdatedAt = time.Now().UTC()

	if input.EndTime != nil {
		t, err := time.Parse(time.RFC3339, *input.EndTime)
		if err != nil {
			return nil, domain.ErrInvalidInput
		}
		event.EndTime = &t
	}

	if err := s.repo.Update(event); err != nil {
		return nil, err
	}
	return event, nil
}
