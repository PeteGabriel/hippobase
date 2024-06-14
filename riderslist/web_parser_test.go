package riderslist

import (
	"slices"
	"testing"
)

func TestParse(t *testing.T) {
	parsedEvents := Parse()
	if len(parsedEvents) == 0 {
		t.Errorf("Expected eventInfo to be greater than 0")
	}
}

func TestScrap_ListOfEventsBiggerThanZero(t *testing.T) {
	// call scrapEventsTable
	events, err := scrapEventsTable(EquinisURL)
	if err != nil {
		t.Errorf("Expected error to be nil")
	}

	assertEventsList(t, events)
}

func TestScrap_ListOfEventsEmpty_IfNoEventsAreFound(t *testing.T) {
	// call scrapEventsTable
	events, err := scrapEventsTable("https://www.example.com")
	if err != nil {
		t.Errorf("Expected error to be nil. Got %v", err)
	}

	if len(events) != 0 {
		t.Errorf("Expected events to be greater than 0")
	}
}

func assertEventsList(t *testing.T, events map[string][]EventEntryRow) {
	if len(events) == 0 {
		t.Errorf("Expected events to be greater than 0")
	}

	ks := []string{"Upcoming", "Recent"}

	for key, value := range events {
		if !slices.Contains(ks, key) {
			t.Errorf("Expected key to be in %v. Got '%v'", ks, key)
		}

		for _, v := range value {
			if len(v.Date) == 0 {
				t.Errorf("Expected date to be greater than 0. Got '%v'", v.Date)
			}
			if len(v.Name) == 0 {
				t.Errorf("Expected name to be greater than 0. Got '%v'", v.Name)
			}
			if len(v.Location) == 0 {
				t.Errorf("Expected location to be greater than 0. Got '%v'", v.Location)
			}
			if len(v.EventURL) == 0 {
				t.Errorf("Expected eventURL to be greater than 0. Got '%v'", v.EventURL)
			}
			if len(v.EntryListURL) == 0 && key != "Upcoming" { // entryListURL is optional for upcoming events
				t.Errorf("Expected entryListURL to be greater than 0. Got '%v'", v.EntryListURL)
			}
		}
	}
}
