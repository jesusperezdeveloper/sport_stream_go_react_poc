package memory

import (
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/jpsdeveloper/sportstream-api/internal/domain"
)

func SeedData(clubRepo *ClubRepository, streamRepo *StreamRepository, eventRepo *EventRepository) {
	now := time.Now().UTC()

	// --- Clubs ---
	lazioID := uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef1234567890")
	sevillaID := uuid.MustParse("b2c3d4e5-f6a7-8901-bcde-f12345678901")
	legaVolleyID := uuid.MustParse("c3d4e5f6-a7b8-9012-cdef-123456789012")
	superTennixID := uuid.MustParse("d4e5f6a7-b8c9-0123-defa-234567890123")
	fibaID := uuid.MustParse("e5f6a7b8-c9d0-1234-efab-345678901234")

	clubs := []domain.Club{
		{
			ID: lazioID, Name: "S.S. Lazio", Slug: "s-s-lazio", Country: "Italy",
			League: "Serie A", LogoURL: "https://images.unsplash.com/photo-1574629810360-7efbbe195018?w=200&h=200&fit=crop",
			Sport: "football", IsActive: true, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: sevillaID, Name: "Sevilla FC", Slug: "sevilla-fc", Country: "Spain",
			League: "La Liga", LogoURL: "https://images.unsplash.com/photo-1522778119026-d647f0596c20?w=200&h=200&fit=crop",
			Sport: "football", IsActive: true, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: legaVolleyID, Name: "Lega Volley", Slug: "lega-volley", Country: "Italy",
			League: "SuperLega", LogoURL: "https://images.unsplash.com/photo-1553778263-73a83bab9b0c?w=200&h=200&fit=crop",
			Sport: "volleyball", IsActive: true, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: superTennixID, Name: "SuperTennix", Slug: "supertennix", Country: "International",
			League: "ATP Tour", LogoURL: "https://images.unsplash.com/photo-1554068865-24cecd4e34b8?w=200&h=200&fit=crop",
			Sport: "tennis", IsActive: true, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: fibaID, Name: "FIBA Europe", Slug: "fiba-europe", Country: "International",
			League: "EuroLeague", LogoURL: "https://images.unsplash.com/photo-1546519638-68e109498ffc?w=200&h=200&fit=crop",
			Sport: "basketball", IsActive: true, CreatedAt: now, UpdatedAt: now,
		},
	}

	for i := range clubs {
		if err := clubRepo.Create(&clubs[i]); err != nil {
			slog.Error("failed to seed club", "name", clubs[i].Name, "error", err)
		}
	}

	// --- Streams ---
	stream1ID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	stream2ID := uuid.MustParse("22222222-2222-2222-2222-222222222222")
	stream3ID := uuid.MustParse("33333333-3333-3333-3333-333333333333")
	stream4ID := uuid.MustParse("44444444-4444-4444-4444-444444444444")
	stream5ID := uuid.MustParse("55555555-5555-5555-5555-555555555555")
	stream6ID := uuid.MustParse("66666666-6666-6666-6666-666666666666")
	stream7ID := uuid.MustParse("77777777-7777-7777-7777-777777777777")
	stream8ID := uuid.MustParse("88888888-8888-8888-8888-888888888888")
	stream9ID := uuid.MustParse("99999999-9999-9999-9999-999999999999")
	stream10ID := uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")

	scheduledTime := now.Add(2 * time.Hour)
	pastStart := now.Add(-1 * time.Hour)
	pastEnd := now.Add(-30 * time.Minute)

	streams := []domain.Stream{
		{
			ID: stream1ID, ClubID: lazioID, Title: "Lazio vs Roma — Serie A Matchday 28",
			Description: "Live coverage of the Rome derby", Type: domain.StreamTypeLive,
			Status: domain.StreamStatusLive, StreamURL: "https://cph-p2p-msl.akamaized.net/hls/live/2000341/test/master.m3u8",
			ThumbnailURL: "https://images.unsplash.com/photo-1574629810360-7efbbe195018?w=400&h=225&fit=crop",
			ViewCount: 320000, Duration: 0, StartedAt: &pastStart,
			Tags: []string{"football", "serie-a", "derby"}, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: stream2ID, ClubID: lazioID, Title: "Lazio Training Session — Behind the Scenes",
			Description: "Exclusive access to the training ground", Type: domain.StreamTypeBehindTheScenes,
			Status: domain.StreamStatusArchived, StreamURL: "https://cdn.jwplayer.com/manifests/pZxWPRg4.m3u8",
			ThumbnailURL: "https://images.unsplash.com/photo-1526232761682-d26e03ac148e?w=400&h=225&fit=crop",
			ViewCount: 45000, Duration: 1800, StartedAt: &pastStart, EndedAt: &pastEnd,
			Tags: []string{"football", "training", "exclusive"}, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: stream3ID, ClubID: sevillaID, Title: "Sevilla vs Betis — La Liga Matchday 30",
			Description: "The Seville derby live from Ramon Sanchez Pizjuan", Type: domain.StreamTypeLive,
			Status: domain.StreamStatusScheduled, StreamURL: "https://cph-p2p-msl.akamaized.net/hls/live/2000341/test/master.m3u8",
			ThumbnailURL: "https://images.unsplash.com/photo-1522778119026-d647f0596c20?w=400&h=225&fit=crop",
			ViewCount: 0, Duration: 0, ScheduledAt: &scheduledTime,
			Tags: []string{"football", "la-liga", "derby"}, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: stream4ID, ClubID: sevillaID, Title: "Sevilla FC — Season Highlights 2025/26",
			Description: "Best moments of the season", Type: domain.StreamTypeHighlight,
			Status: domain.StreamStatusArchived, StreamURL: "https://demo.unified-streaming.com/k8s/features/stable/video/tears-of-steel/tears-of-steel.ism/.m3u8",
			ThumbnailURL: "https://images.unsplash.com/photo-1459865264687-595d652de67e?w=400&h=225&fit=crop",
			ViewCount: 150000, Duration: 720, StartedAt: &pastStart, EndedAt: &pastEnd,
			Tags: []string{"football", "highlights"}, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: stream5ID, ClubID: legaVolleyID, Title: "Perugia vs Trento — SuperLega Final",
			Description: "SuperLega championship final live", Type: domain.StreamTypeLive,
			Status: domain.StreamStatusLive, StreamURL: "https://cph-p2p-msl.akamaized.net/hls/live/2000341/test/master.m3u8",
			ThumbnailURL: "https://images.unsplash.com/photo-1553778263-73a83bab9b0c?w=400&h=225&fit=crop",
			ViewCount: 89000, Duration: 0, StartedAt: &pastStart,
			Tags: []string{"volleyball", "superlega", "final"}, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: stream6ID, ClubID: legaVolleyID, Title: "SuperLega — Best Rallies of the Week",
			Description: "Top rallies compilation", Type: domain.StreamTypeHighlight,
			Status: domain.StreamStatusArchived, StreamURL: "https://storage.googleapis.com/shaka-demo-assets/angel-one-hls/hls.m3u8",
			ThumbnailURL: "https://images.unsplash.com/photo-1530549387789-4c1017266635?w=400&h=225&fit=crop",
			ViewCount: 62000, Duration: 600, StartedAt: &pastStart, EndedAt: &pastEnd,
			Tags: []string{"volleyball", "highlights"}, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: stream7ID, ClubID: superTennixID, Title: "ATP Rome Masters — Sinner vs Djokovic",
			Description: "Semi-final live from Foro Italico", Type: domain.StreamTypeLive,
			Status: domain.StreamStatusEnded, StreamURL: "https://storage.googleapis.com/shaka-demo-assets/bbb-dark-truths-hls/hls.m3u8",
			ThumbnailURL: "https://images.unsplash.com/photo-1554068865-24cecd4e34b8?w=400&h=225&fit=crop",
			ViewCount: 510000, Duration: 7200, StartedAt: &pastStart, EndedAt: &pastEnd,
			Tags: []string{"tennis", "atp", "masters"}, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: stream8ID, ClubID: superTennixID, Title: "Roland Garros — Day 5 Highlights",
			Description: "Best matches from day 5", Type: domain.StreamTypeVOD,
			Status: domain.StreamStatusArchived, StreamURL: "https://devstreaming-cdn.apple.com/videos/streaming/examples/bipbop_16x9/bipbop_16x9_variant.m3u8",
			ThumbnailURL: "https://images.unsplash.com/photo-1504016798967-59a258e9386d?w=400&h=225&fit=crop",
			ViewCount: 230000, Duration: 3600, StartedAt: &pastStart, EndedAt: &pastEnd,
			Tags: []string{"tennis", "grand-slam", "vod"}, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: stream9ID, ClubID: fibaID, Title: "EuroLeague — Real Madrid vs Olympiacos",
			Description: "Quarter-final game 3", Type: domain.StreamTypeLive,
			Status: domain.StreamStatusScheduled, StreamURL: "https://cph-p2p-msl.akamaized.net/hls/live/2000341/test/master.m3u8",
			ThumbnailURL: "https://images.unsplash.com/photo-1546519638-68e109498ffc?w=400&h=225&fit=crop",
			ViewCount: 0, Duration: 0, ScheduledAt: &scheduledTime,
			Tags: []string{"basketball", "euroleague", "playoffs"}, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: stream10ID, ClubID: fibaID, Title: "FIBA Europe — Top 10 Plays March 2026",
			Description: "Monthly top plays compilation", Type: domain.StreamTypeHighlight,
			Status: domain.StreamStatusArchived, StreamURL: "https://cdn.jwplayer.com/manifests/pZxWPRg4.m3u8",
			ThumbnailURL: "https://images.unsplash.com/photo-1471295253337-3ceaaedca402?w=400&h=225&fit=crop",
			ViewCount: 175000, Duration: 480, StartedAt: &pastStart, EndedAt: &pastEnd,
			Tags: []string{"basketball", "highlights", "top-plays"}, CreatedAt: now, UpdatedAt: now,
		},
	}

	for i := range streams {
		if err := streamRepo.Create(&streams[i]); err != nil {
			slog.Error("failed to seed stream", "title", streams[i].Title, "error", err)
		}
	}

	// --- Events ---
	event1Start := now.Add(3 * time.Hour)
	event2Start := now.Add(24 * time.Hour)
	event3End := now.Add(-15 * time.Minute)
	event4Start := now.Add(48 * time.Hour)
	event5Start := now.Add(72 * time.Hour)

	events := []domain.Event{
		{
			ID: uuid.MustParse("f1111111-1111-1111-1111-111111111111"),
			ClubID: lazioID, Title: "Lazio vs Roma — Serie A Matchday 28",
			Description: "Rome derby at Stadio Olimpico", Venue: "Stadio Olimpico, Rome",
			Sport: "football", StartTime: pastStart, Status: domain.EventStatusLive,
			StreamID: &stream1ID, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uuid.MustParse("f2222222-2222-2222-2222-222222222222"),
			ClubID: sevillaID, Title: "Sevilla vs Betis — La Liga Matchday 30",
			Description: "The great Seville derby", Venue: "Ramon Sanchez Pizjuan, Seville",
			Sport: "football", StartTime: event1Start, Status: domain.EventStatusUpcoming,
			StreamID: &stream3ID, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uuid.MustParse("f3333333-3333-3333-3333-333333333333"),
			ClubID: superTennixID, Title: "ATP Rome Masters — Semi-finals",
			Description: "Semi-final day at Foro Italico", Venue: "Foro Italico, Rome",
			Sport: "tennis", StartTime: pastStart, EndTime: &event3End,
			Status: domain.EventStatusCompleted, StreamID: &stream7ID,
			CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uuid.MustParse("f4444444-4444-4444-4444-444444444444"),
			ClubID: legaVolleyID, Title: "SuperLega Final — Game 2",
			Description: "Championship final second leg", Venue: "PalaBarton, Perugia",
			Sport: "volleyball", StartTime: event2Start, Status: domain.EventStatusUpcoming,
			CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uuid.MustParse("f5555555-5555-5555-5555-555555555555"),
			ClubID: fibaID, Title: "EuroLeague — Real Madrid vs Olympiacos Game 3",
			Description: "Quarter-final game 3", Venue: "WiZink Center, Madrid",
			Sport: "basketball", StartTime: event4Start, Status: domain.EventStatusUpcoming,
			StreamID: &stream9ID, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uuid.MustParse("f6666666-6666-6666-6666-666666666666"),
			ClubID: lazioID, Title: "Lazio vs Napoli — Serie A Matchday 29",
			Description: "Key match in the title race", Venue: "Stadio Olimpico, Rome",
			Sport: "football", StartTime: event5Start, Status: domain.EventStatusUpcoming,
			CreatedAt: now, UpdatedAt: now,
		},
	}

	for i := range events {
		if err := eventRepo.Create(&events[i]); err != nil {
			slog.Error("failed to seed event", "title", events[i].Title, "error", err)
		}
	}

	slog.Info("seed data loaded", "clubs", len(clubs), "streams", len(streams), "events", len(events))
}
