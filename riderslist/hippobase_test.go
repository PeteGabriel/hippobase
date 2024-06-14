package riderslist

import "testing"

func TestGetEntryLists_ForACorrectTableOfEvents(t *testing.T) {

	// Arrange
	eventsTable := RelatedDateEventsTable{
		"Upcoming": []EventEntryRow{
			{
				EntryListURL: "https://www.hippobase.com/EventInfo/Entries/CompetitorHorse.aspx?EventID=858&lang=en&EntryGroupIDs=CSIO&Submit=OK&ShowCompetitors=on&ShowHorses=on&HorsesCompact=on",
			},
		},
	}

	// Act
	eventInfo := GetEntryLists(eventsTable)

	if len(eventInfo) == 0 {
		t.Error("Expected eventInfo to be greater than 0")
	}

	if eventInfo[0].TotalAthletes != len(eventInfo[0].Competitors) {
		t.Error("Total athlete must match total of competitors - ", eventInfo[0].TotalAthletes, "!=", len(eventInfo[0].Competitors))
	}
}
