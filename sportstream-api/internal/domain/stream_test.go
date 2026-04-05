package domain

import "testing"

func TestValidStreamType(t *testing.T) {
	validTypes := []string{"live", "vod", "highlight", "behind_the_scenes"}
	for _, st := range validTypes {
		if !IsValidStreamType(st) {
			t.Errorf("expected %q to be a valid stream type", st)
		}
	}
}

func TestInvalidStreamType(t *testing.T) {
	invalidTypes := []string{"", "podcast", "LIVE", "Live", "unknown", "replay"}
	for _, st := range invalidTypes {
		if IsValidStreamType(st) {
			t.Errorf("expected %q to be an invalid stream type", st)
		}
	}
}

func TestValidStreamStatus(t *testing.T) {
	validStatuses := []string{"scheduled", "live", "ended", "archived"}
	for _, s := range validStatuses {
		if !IsValidStreamStatus(s) {
			t.Errorf("expected %q to be a valid stream status", s)
		}
	}
}

func TestInvalidStreamStatus(t *testing.T) {
	invalidStatuses := []string{"", "paused", "LIVE", "Live", "cancelled", "deleted"}
	for _, s := range invalidStatuses {
		if IsValidStreamStatus(s) {
			t.Errorf("expected %q to be an invalid stream status", s)
		}
	}
}

func TestValidStatusTransition(t *testing.T) {
	transitions := []struct {
		from StreamStatus
		to   StreamStatus
	}{
		{StreamStatusScheduled, StreamStatusLive},
		{StreamStatusLive, StreamStatusEnded},
		{StreamStatusEnded, StreamStatusArchived},
	}

	for _, tr := range transitions {
		if !IsValidStreamTransition(tr.from, tr.to) {
			t.Errorf("expected transition %q -> %q to be valid", tr.from, tr.to)
		}
	}
}

func TestInvalidStatusTransition(t *testing.T) {
	transitions := []struct {
		from StreamStatus
		to   StreamStatus
	}{
		{StreamStatusLive, StreamStatusScheduled},
		{StreamStatusEnded, StreamStatusLive},
		{StreamStatusScheduled, StreamStatusEnded},
		{StreamStatusScheduled, StreamStatusArchived},
		{StreamStatusLive, StreamStatusArchived},
		{StreamStatusArchived, StreamStatusLive},
		{StreamStatusArchived, StreamStatusScheduled},
		{StreamStatusArchived, StreamStatusEnded},
	}

	for _, tr := range transitions {
		if IsValidStreamTransition(tr.from, tr.to) {
			t.Errorf("expected transition %q -> %q to be invalid", tr.from, tr.to)
		}
	}
}

func TestValidStreamTypesReturnsAll(t *testing.T) {
	types := ValidStreamTypes()
	if len(types) != 4 {
		t.Errorf("expected 4 valid stream types, got %d", len(types))
	}
}

func TestValidStreamStatusesReturnsAll(t *testing.T) {
	statuses := ValidStreamStatuses()
	if len(statuses) != 4 {
		t.Errorf("expected 4 valid stream statuses, got %d", len(statuses))
	}
}

func TestIsValidStreamTransitionUnknownFrom(t *testing.T) {
	if IsValidStreamTransition("unknown", StreamStatusLive) {
		t.Error("expected transition from unknown status to be invalid")
	}
}
