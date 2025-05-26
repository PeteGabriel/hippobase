package hippobase

import (
	"testing"
)

func TestGetEvents_Success(t *testing.T) {
	events, err := GetEvents()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(events) == 0 {
		t.Fatalf("Expected events to be greater than 0")
	}
}

func TestGetEntryLists(t *testing.T) {
	events, err := GetEvents()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	url := events.FirstEntryListURL()
	competition, err := GetEntryLists(url)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if competition == nil {
		t.Fatalf("Expected competition to be not nil")
	}
	if len(competition.Events) == 0 {
		t.Fatalf("Expected competitors to be greater than 0")
	}

	if competition.MainTitle == "" {
		t.Fatalf("Expected MainTitle to be not empty")
	}
}
