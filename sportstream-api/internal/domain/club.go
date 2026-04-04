package domain

import (
	"time"

	"github.com/google/uuid"
)

type Club struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Country   string    `json:"country"`
	League    string    `json:"league"`
	LogoURL   string    `json:"logo_url"`
	Sport     string    `json:"sport"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ClubRepository interface {
	FindAll() ([]Club, error)
	FindByID(id uuid.UUID) (*Club, error)
	Create(club *Club) error
	Update(club *Club) error
	Delete(id uuid.UUID) error
}
