package memory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/jpsdeveloper/sportstream-api/internal/domain"
)

type ClubRepository struct {
	mu    sync.RWMutex
	clubs map[uuid.UUID]*domain.Club
}

func NewClubRepository() *ClubRepository {
	return &ClubRepository{
		clubs: make(map[uuid.UUID]*domain.Club),
	}
}

func (r *ClubRepository) FindAll() ([]domain.Club, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	clubs := make([]domain.Club, 0, len(r.clubs))
	for _, c := range r.clubs {
		clubs = append(clubs, *c)
	}
	return clubs, nil
}

func (r *ClubRepository) FindByID(id uuid.UUID) (*domain.Club, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	club, ok := r.clubs[id]
	if !ok {
		return nil, domain.ErrClubNotFound
	}
	copy := *club
	return &copy, nil
}

func (r *ClubRepository) Create(club *domain.Club) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	copy := *club
	r.clubs[club.ID] = &copy
	return nil
}

func (r *ClubRepository) Update(club *domain.Club) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.clubs[club.ID]; !ok {
		return domain.ErrClubNotFound
	}
	copy := *club
	r.clubs[club.ID] = &copy
	return nil
}

func (r *ClubRepository) Delete(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.clubs[id]; !ok {
		return domain.ErrClubNotFound
	}
	delete(r.clubs, id)
	return nil
}
