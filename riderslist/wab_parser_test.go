package riderslist

import (
	"slices"
	"testing"
)

// test Parse
func TestParse(t *testing.T) {
	// call Parse
	Parse()
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
			if len(v.date) == 0 {
				t.Errorf("Expected date to be greater than 0. Got '%v'", v.date)
			}
			if len(v.name) == 0 {
				t.Errorf("Expected name to be greater than 0. Got '%v'", v.name)
			}
			if len(v.location) == 0 {
				t.Errorf("Expected location to be greater than 0. Got '%v'", v.location)
			}
			if len(v.eventURL) == 0 {
				t.Errorf("Expected eventURL to be greater than 0. Got '%v'", v.eventURL)
			}
			if len(v.entryListURL) == 0 && key != "Upcoming" { // entryListURL is optional for upcoming events
				t.Errorf("Expected entryListURL to be greater than 0. Got '%v'", v.entryListURL)
			}
		}
	}
}
