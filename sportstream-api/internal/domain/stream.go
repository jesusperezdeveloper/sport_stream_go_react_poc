package domain

import (
	"time"

	"github.com/google/uuid"
)

type StreamType string

const (
	StreamTypeLive             StreamType = "live"
	StreamTypeVOD              StreamType = "vod"
	StreamTypeHighlight        StreamType = "highlight"
	StreamTypeBehindTheScenes  StreamType = "behind_the_scenes"
)

func ValidStreamTypes() []StreamType {
	return []StreamType{StreamTypeLive, StreamTypeVOD, StreamTypeHighlight, StreamTypeBehindTheScenes}
}

func IsValidStreamType(t string) bool {
	for _, valid := range ValidStreamTypes() {
		if string(valid) == t {
			return true
		}
	}
	return false
}

type StreamStatus string

const (
	StreamStatusScheduled StreamStatus = "scheduled"
	StreamStatusLive      StreamStatus = "live"
	StreamStatusEnded     StreamStatus = "ended"
	StreamStatusArchived  StreamStatus = "archived"
)

func ValidStreamStatuses() []StreamStatus {
	return []StreamStatus{StreamStatusScheduled, StreamStatusLive, StreamStatusEnded, StreamStatusArchived}
}

func IsValidStreamStatus(s string) bool {
	for _, valid := range ValidStreamStatuses() {
		if string(valid) == s {
			return true
		}
	}
	return false
}

var validStreamTransitions = map[StreamStatus][]StreamStatus{
	StreamStatusScheduled: {StreamStatusLive},
	StreamStatusLive:      {StreamStatusEnded},
	StreamStatusEnded:     {StreamStatusArchived},
	StreamStatusArchived:  {},
}

func IsValidStreamTransition(from, to StreamStatus) bool {
	allowed, ok := validStreamTransitions[from]
	if !ok {
		return false
	}
	for _, s := range allowed {
		if s == to {
			return true
		}
	}
	return false
}

type Stream struct {
	ID           uuid.UUID    `json:"id"`
	ClubID       uuid.UUID    `json:"club_id"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	Type         StreamType   `json:"type"`
	Status       StreamStatus `json:"status"`
	StreamURL    string       `json:"stream_url"`
	ThumbnailURL string       `json:"thumbnail_url"`
	ViewCount    int64        `json:"view_count"`
	Duration     int          `json:"duration"`
	ScheduledAt  *time.Time   `json:"scheduled_at,omitempty"`
	StartedAt    *time.Time   `json:"started_at,omitempty"`
	EndedAt      *time.Time   `json:"ended_at,omitempty"`
	Tags         []string     `json:"tags"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

type StreamRepository interface {
	FindAll() ([]Stream, error)
	FindByID(id uuid.UUID) (*Stream, error)
	FindByClubID(clubID uuid.UUID) ([]Stream, error)
	Create(stream *Stream) error
	Update(stream *Stream) error
	UpdateStatus(id uuid.UUID, status StreamStatus) error
}
