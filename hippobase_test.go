package equestrian_events_riders_list

import "testing"

func TestGetEntryLists_ForACorrectTableOfEvents(t *testing.T) {

	// Arrange
	eventsTable := RelatedDateEventsTable{
		"Upcoming": []EventEntryRow{
			{
				EntryListURL: "https://www.hippobase.com/EventInfo/Entries/CompetitorHorse.aspx?EventID=863",
			},
			{
				EntryListURL: "https://www.hippobase.com/EventInfo/Entries/CompetitorHorse.aspx?EventID=873",
			},
		},
		"Recent": []EventEntryRow{
			{
				EntryListURL: "https://www.hippobase.com/EventInfo/Entries/CompetitorHorse.aspx?EventID=858",
			},
			{
				EntryListURL: "https://www.hippobase.com/EventInfo/Entries/CompetitorHorse.aspx?EventID=860",
			},
		},
	}

	// Act
	competitions := GetEntryLists(eventsTable)

	if len(competitions) != 4 {
		t.Error("Expected competitions to be eql to 4")
	}

	if len(competitions[0].Events) == 0 {
		t.Error("Expected events to be greater than 0")
	}

	if len(competitions[1].Events) == 0 {
		t.Error("Expected events to be greater than 0")
	}

	if len(competitions[2].Events) == 0 {
		t.Error("Expected events to be greater than 0")
	}

	if len(competitions[3].Events) == 0 {
		t.Error("Expected events to be greater than 0")
	}

}
