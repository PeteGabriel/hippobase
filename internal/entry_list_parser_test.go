package internal

import (
	"testing"
)

func TestParseEvent(t *testing.T) {
	entryList := "https://www.hippobase.com/EventInfo/Entries/CompetitorHorse.aspx?EventID=850"
	competition, err := ParseEvent(entryList)

	if err != nil {
		t.Fatalf("Error parsing event: %v", err)
	}

	if competition == nil {
		t.Fatalf("Expected competition to be not nil")
	}

	if competition.MainTitle == "" {
		t.Fatalf("Expected MainTitle to be not empty")
	}

	if competition.Events == nil {
		t.Fatalf("Expected Events to be not nil")
	}

	if len(competition.Events) == 0 {
		t.Fatalf("Expected Events to be greater than 0")
	}
	for _, event := range competition.Events {

		for _, competitor := range event.Competitors {
			if competitor.Flag == "" {
				t.Fatalf("Expected Flag to be not empty")
			}
			if competitor.CountryCode == "" {
				t.Fatalf("Expected CountryCode to be not empty")
			}
			if competitor.CountryName == "" {
				t.Fatalf("Expected CountryName to be not empty")
			}
			for _, pair := range competitor.Pairs {
				if pair.Competitor == "" {
					t.Fatalf("Expected Competitor to be not empty")
				}
				if len(pair.Horses) == 0 {
					t.Fatalf("Expected Horses to be greater than 0")
				}
			}
		}

		if event.TotalNations == 0 {
			t.Fatalf("Expected TotalNations to be greater than 0")
		}
		if event.EventFullName == "" {
			t.Fatalf("Expected EventFullName to be not empty")
		}
		if event.CreatedAt == "" {
			t.Fatalf("Expected CreatedAt to be not empty")
		}
		if event.TotalAthletes == 0 {
			t.Fatalf("Expected TotalAthletes to be greater than 0")
		}
		if event.TotalHorses == 0 {
			t.Fatalf("Expected TotalHorses to be greater than 0")
		}
	}
}
