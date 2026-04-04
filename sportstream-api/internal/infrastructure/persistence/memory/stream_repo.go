package memory

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/jpsdeveloper/sportstream-api/internal/domain"
)

type StreamRepository struct {
	mu      sync.RWMutex
	streams map[uuid.UUID]*domain.Stream
}

func NewStreamRepository() *StreamRepository {
	return &StreamRepository{
		streams: make(map[uuid.UUID]*domain.Stream),
	}
}

func (r *StreamRepository) FindAll() ([]domain.Stream, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	streams := make([]domain.Stream, 0, len(r.streams))
	for _, s := range r.streams {
		streams = append(streams, *s)
	}
	return streams, nil
}

func (r *StreamRepository) FindByID(id uuid.UUID) (*domain.Stream, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	stream, ok := r.streams[id]
	if !ok {
		return nil, domain.ErrStreamNotFound
	}
	copy := *stream
	return &copy, nil
}

func (r *StreamRepository) FindByClubID(clubID uuid.UUID) ([]domain.Stream, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var streams []domain.Stream
	for _, s := range r.streams {
		if s.ClubID == clubID {
			streams = append(streams, *s)
		}
	}
	return streams, nil
}

func (r *StreamRepository) Create(stream *domain.Stream) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	copy := *stream
	if copy.Tags == nil {
		copy.Tags = []string{}
	}
	r.streams[stream.ID] = &copy
	return nil
}

func (r *StreamRepository) Update(stream *domain.Stream) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.streams[stream.ID]; !ok {
		return domain.ErrStreamNotFound
	}
	copy := *stream
	r.streams[stream.ID] = &copy
	return nil
}

func (r *StreamRepository) UpdateStatus(id uuid.UUID, status domain.StreamStatus) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	stream, ok := r.streams[id]
	if !ok {
		return domain.ErrStreamNotFound
	}

	if !domain.IsValidStreamTransition(stream.Status, status) {
		return domain.ErrInvalidTransition
	}

	stream.Status = status
	now := time.Now().UTC()
	stream.UpdatedAt = now

	if status == domain.StreamStatusLive {
		stream.StartedAt = &now
	} else if status == domain.StreamStatusEnded {
		stream.EndedAt = &now
	}

	return nil
}
