package dashboard

import (
	"sort"

	"github.com/google/uuid"
	"github.com/jpsdeveloper/sportstream-api/internal/domain"
)

type Service struct {
	clubRepo   domain.ClubRepository
	streamRepo domain.StreamRepository
	eventRepo  domain.EventRepository
}

func NewService(clubRepo domain.ClubRepository, streamRepo domain.StreamRepository, eventRepo domain.EventRepository) *Service {
	return &Service{
		clubRepo:   clubRepo,
		streamRepo: streamRepo,
		eventRepo:  eventRepo,
	}
}

type ClubViewSummary struct {
	ClubID     uuid.UUID `json:"club_id"`
	ClubName   string    `json:"club_name"`
	TotalViews int64     `json:"total_views"`
}

type Summary struct {
	TotalClubs      int                `json:"total_clubs"`
	TotalStreams     int                `json:"total_streams"`
	LiveStreams      int                `json:"live_streams"`
	UpcomingEvents   int                `json:"upcoming_events"`
	TotalViews      int64              `json:"total_views"`
	StreamsByType   map[string]int     `json:"streams_by_type"`
	StreamsByStatus map[string]int     `json:"streams_by_status"`
	TopClubsByViews []ClubViewSummary  `json:"top_clubs_by_views"`
}

func (s *Service) GetSummary() (*Summary, error) {
	clubs, err := s.clubRepo.FindAll()
	if err != nil {
		return nil, err
	}

	streams, err := s.streamRepo.FindAll()
	if err != nil {
		return nil, err
	}

	events, err := s.eventRepo.FindAll()
	if err != nil {
		return nil, err
	}

	summary := &Summary{
		TotalClubs:      len(clubs),
		TotalStreams:     len(streams),
		StreamsByType:   make(map[string]int),
		StreamsByStatus: make(map[string]int),
	}

	clubMap := make(map[uuid.UUID]string)
	for _, c := range clubs {
		clubMap[c.ID] = c.Name
	}

	clubViews := make(map[uuid.UUID]int64)
	for _, st := range streams {
		summary.TotalViews += st.ViewCount
		summary.StreamsByType[string(st.Type)]++
		summary.StreamsByStatus[string(st.Status)]++
		if st.Status == domain.StreamStatusLive {
			summary.LiveStreams++
		}
		clubViews[st.ClubID] += st.ViewCount
	}

	for _, e := range events {
		if e.Status == domain.EventStatusUpcoming {
			summary.UpcomingEvents++
		}
	}

	topClubs := make([]ClubViewSummary, 0, len(clubViews))
	for clubID, views := range clubViews {
		topClubs = append(topClubs, ClubViewSummary{
			ClubID:     clubID,
			ClubName:   clubMap[clubID],
			TotalViews: views,
		})
	}
	sort.Slice(topClubs, func(i, j int) bool {
		return topClubs[i].TotalViews > topClubs[j].TotalViews
	})
	summary.TopClubsByViews = topClubs

	return summary, nil
}
