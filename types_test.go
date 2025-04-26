package hippobase

import (
	"encoding/json"
	"testing"
)

// TestEntryRow tests the RidersEntryRow struct for JSON marshaling and unmarshaling.
func TestEntryRow(t *testing.T) {
	entry := RidersEntryRow{
		Flag:        "flag.png",
		CountryCode: "USA",
		CountryName: "United States",
		Pairs: map[string][]string{
			"John Doe":   {"Horse1", "Horse2"},
			"Jane Smith": {"Horse3"},
		},
	}

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
