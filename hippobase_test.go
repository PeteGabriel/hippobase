package equestrian_events_riders_list

import "testing"

func TestGetEntryLists_ForACorrectTableOfEvents(t *testing.T) {

	// Arrange
	eventsTable := RelatedDateEventsTable{
		"Upcoming": []EventEntryRow{
			{
				EntryListURL: "https://www.hippobase.com/EventInfo/Entries/CompetitorHorse.aspx?EventID=858",
			},
		},
	}

	// Act
	eventInfo := GetEntryLists(eventsTable)

	if len(eventInfo) != 1 {
		t.Error("Expected eventInfo to be eql to 1")
	}

	if len(eventInfo[0].events) != 2 {
		t.Error("Expected events to be eql to 2")
	}

	if eventInfo[0].MainTitle != "CSIO St. Gallen 2024" {
		t.Error("Main title must match CSIO St. Gallen 2024")
	}

	if eventInfo[0].events[0].EventFullName != "Entries CSIO" {
		t.Error("Event name must match Entries CSIO")
	}
	if eventInfo[0].events[1].EventFullName != "Entries CSN" {
		t.Error("Event name must match Entries CSN")
	}

	if eventInfo[0].events[0].TotalNations != 15 {
		t.Error("Total nations must match total of 15")
	}
	if eventInfo[0].events[0].TotalAthletes != 73 {
		t.Error("Total athletes must match total of 73")
	}
	if eventInfo[0].events[0].TotalHorses != 225 {
		t.Error("Total horses must match total of 225")
	}
}
