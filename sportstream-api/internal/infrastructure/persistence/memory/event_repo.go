package memory

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/jpsdeveloper/sportstream-api/internal/domain"
)

type EventRepository struct {
	mu     sync.RWMutex
	events map[uuid.UUID]*domain.Event
}

func NewEventRepository() *EventRepository {
	return &EventRepository{
		events: make(map[uuid.UUID]*domain.Event),
	}
}

func (r *EventRepository) FindAll() ([]domain.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	events := make([]domain.Event, 0, len(r.events))
	for _, e := range r.events {
		events = append(events, *e)
	}
	return events, nil
}

func (r *EventRepository) FindByID(id uuid.UUID) (*domain.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	event, ok := r.events[id]
	if !ok {
		return nil, domain.ErrEventNotFound
	}
	copy := *event
	return &copy, nil
}

func (r *EventRepository) FindUpcoming() ([]domain.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	now := time.Now().UTC()
	var events []domain.Event
	for _, e := range r.events {
		if e.Status == domain.EventStatusUpcoming && e.StartTime.After(now) {
			events = append(events, *e)
		}
	}
	return events, nil
}

func (r *EventRepository) Create(event *domain.Event) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	copy := *event
	r.events[event.ID] = &copy
	return nil
}

func (r *EventRepository) Update(event *domain.Event) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.events[event.ID]; !ok {
		return domain.ErrEventNotFound
	}
	copy := *event
	r.events[event.ID] = &copy
	return nil
}
