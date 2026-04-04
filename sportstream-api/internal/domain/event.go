package domain

import (
	"time"

	"github.com/google/uuid"
)

type EventStatus string

const (
	EventStatusUpcoming  EventStatus = "upcoming"
	EventStatusLive      EventStatus = "live"
	EventStatusCompleted EventStatus = "completed"
	EventStatusCancelled EventStatus = "cancelled"
)

func ValidEventStatuses() []EventStatus {
	return []EventStatus{EventStatusUpcoming, EventStatusLive, EventStatusCompleted, EventStatusCancelled}
}

func IsValidEventStatus(s string) bool {
	for _, valid := range ValidEventStatuses() {
		if string(valid) == s {
			return true
		}
	}
	return false
}

type Event struct {
	ID          uuid.UUID   `json:"id"`
	ClubID      uuid.UUID   `json:"club_id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Venue       string      `json:"venue"`
	Sport       string      `json:"sport"`
	StartTime   time.Time   `json:"start_time"`
	EndTime     *time.Time  `json:"end_time,omitempty"`
	Status      EventStatus `json:"status"`
	StreamID    *uuid.UUID  `json:"stream_id,omitempty"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type EventRepository interface {
	FindAll() ([]Event, error)
	FindByID(id uuid.UUID) (*Event, error)
	FindUpcoming() ([]Event, error)
	Create(event *Event) error
	Update(event *Event) error
}
