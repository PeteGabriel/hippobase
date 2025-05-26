package internal

import (
	"encoding/json"
	"testing"
)

func getEntry() *RidersEntryRow {
	return &RidersEntryRow{
		Flag:        "flag.png",
		CountryCode: "USA",
		CountryName: "United States",
		Pairs: map[string][]string{
			"John Doe":   {"Horse1", "Horse2"},
			"Jane Smith": {"Horse3"},
		},
	}
}

// TestEntryRow tests the RidersEntryRow struct for JSON marshaling and unmarshaling.
func TestEntryRow(t *testing.T) {
	entry := getEntry()

	marshaled, err := json.Marshal(entry)
	if err != nil {
		t.Fatalf("Failed to marshal RidersEntryRow: %v", err)
	}

	// unmarshal
	var unmarshaled RidersEntryRow
	err = json.Unmarshal(marshaled, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal RidersEntryRow: %v", err)
	}

	// compare fields
	if unmarshaled.Flag != entry.Flag {
		t.Fatalf("Expected Flag %s, got %s", entry.Flag, unmarshaled.Flag)
	}
	if unmarshaled.CountryCode != entry.CountryCode {
		t.Fatalf("Expected CountryCode %s, got %s", entry.CountryCode, unmarshaled.CountryCode)
	}
	if unmarshaled.CountryName != entry.CountryName {
		t.Fatalf("Expected CountryName %s, got %s", entry.CountryName, unmarshaled.CountryName)
	}
	if len(unmarshaled.Pairs) != len(entry.Pairs) {
		t.Fatalf("Expected Pairs length %d, got %d", len(entry.Pairs), len(unmarshaled.Pairs))
	}
	for k, v := range entry.Pairs {
		if len(unmarshaled.Pairs[k]) != len(v) {
			t.Fatalf("Expected Pairs[%s] length %d, got %d", k, len(v), len(unmarshaled.Pairs[k]))
		}
		for i, horse := range v {
			if unmarshaled.Pairs[k][i] != horse {
				t.Fatalf("Expected Pairs[%s][%d] %s, got %s", k, i, horse, unmarshaled.Pairs[k][i])
			}
		}
	}
}

func TestEventEntryRow(t *testing.T) {
	entry := EventEntryRow{
		Date:         "2023-10-01",
		Name:         "Equestrian Event",
		Location:     "Location Name",
		EventURL:     "https://example.com/event",
		EntryListURL: "https://example.com/entrylist",
	}

	marshaled, err := json.Marshal(entry)
	if err != nil {
		t.Fatalf("Failed to marshal EventEntryRow: %v", err)
	}

	var unmarshaled EventEntryRow
	err = json.Unmarshal(marshaled, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal EventEntryRow: %v", err)
	}

	if unmarshaled.Date != entry.Date {
		t.Fatalf("Expected Date %s, got %s", entry.Date, unmarshaled.Date)
	}
	if unmarshaled.Name != entry.Name {
		t.Fatalf("Expected Name %s, got %s", entry.Name, unmarshaled.Name)
	}
	if unmarshaled.Location != entry.Location {
		t.Fatalf("Expected Location %s, got %s", entry.Location, unmarshaled.Location)
	}
	if unmarshaled.EventURL != entry.EventURL {
		t.Fatalf("Expected EventURL %s, got %s", entry.EventURL, unmarshaled.EventURL)
	}
	if unmarshaled.EntryListURL != entry.EntryListURL {
		t.Fatalf("Expected EntryListURL %s, got %s", entry.EntryListURL, unmarshaled.EntryListURL)
	}
}

func TestEventInfo(t *testing.T) {
	event := EventInfo{
		CreatedAt:     "2023-10-01",
		EventFullName: "Equestrian Event",
		Competitors: []*RidersEntryRow{
			getEntry(),
		},
		TotalNations:  5,
		TotalAthletes: 10,
		TotalHorses:   15,
	}

	marshaled, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("Failed to marshal EventInfo: %v", err)
	}

	var unmarshaled EventInfo
	err = json.Unmarshal(marshaled, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal EventInfo: %v", err)
	}

	if unmarshaled.CreatedAt != event.CreatedAt {
		t.Fatalf("Expected CreatedAt %s, got %s", event.CreatedAt, unmarshaled.CreatedAt)
	}
	if unmarshaled.EventFullName != event.EventFullName {
		t.Fatalf("Expected EventFullName %s, got %s", event.EventFullName, unmarshaled.EventFullName)
	}
	if len(unmarshaled.Competitors) != len(event.Competitors) {
		t.Fatalf("Expected Competitors length %d, got %d", len(event.Competitors), len(unmarshaled.Competitors))
	}
	for i, competitor := range event.Competitors {
		if unmarshaled.Competitors[i].Flag != competitor.Flag {
			t.Fatalf("Expected Competitors[%d].Flag %s, got %s", i, competitor.Flag, unmarshaled.Competitors[i].Flag)
		}
		if unmarshaled.Competitors[i].CountryCode != competitor.CountryCode {
			t.Fatalf("Expected Competitors[%d].CountryCode %s, got %s", i, competitor.CountryCode, unmarshaled.Competitors[i].CountryCode)
		}
		if unmarshaled.Competitors[i].CountryName != competitor.CountryName {
			t.Fatalf("Expected Competitors[%d].CountryName %s, got %s", i, competitor.CountryName, unmarshaled.Competitors[i].CountryName)
		}
	}
	if unmarshaled.TotalNations != event.TotalNations {
		t.Fatalf("Expected TotalNations %d, got %d", event.TotalNations, unmarshaled.TotalNations)
	}
	if unmarshaled.TotalAthletes != event.TotalAthletes {
		t.Fatalf("Expected TotalAthletes %d, got %d", event.TotalAthletes, unmarshaled.TotalAthletes)
	}
	if unmarshaled.TotalHorses != event.TotalHorses {
		t.Fatalf("Expected TotalHorses %d, got %d", event.TotalHorses, unmarshaled.TotalHorses)
	}
}
