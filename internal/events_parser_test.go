package internal

import (
	"testing"
)

func TestScrapMainTable(t *testing.T) {
	events, err := scrapMainTable(equinis)
	if err != nil {
		t.Fatalf("Error scraping main table: %v", err)
	}

	if len(events) == 0 {
		t.Fatalf("Expected events to be greater than 0")
	}

	// Pick the last event. Highest probability of being the most updated
	lastEvent := events[len(events)-1]
	if lastEvent.EventURL == "" {
		t.Fatalf("Expected EventURL to be present")
	}
	if lastEvent.EntryListURL == "" {
		t.Fatalf("Expected EntryListURL to be present")
	}
	if lastEvent.Date == "" {
		t.Fatalf("Expected Date to be present")
	}
	if lastEvent.Name == "" {
		t.Fatalf("Expected EventFullName to be present")
	}
	if lastEvent.EventURL == "" {
		t.Fatalf("Expected EventURL to be present")
	}
}
