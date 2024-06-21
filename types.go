package equestrian_events_riders_list

// RidersEntryRow represents a row in the entry list of riders for a certain event.
type RidersEntryRow struct {
	Flag, CountryCode, CountryName string
	// horse rider - set of horses
	Pairs map[string][]string
}

// EventEntryRow represents a row in the minified event table.
type EventEntryRow struct {
	Date, Name, Location, EventURL, EntryListURL string
}

// EventInfo represents the complete information of an event.
// It contains the event name, the date of creation and the entry list of competitors.
type EventInfo struct {
	CreatedAt,
	EventFullName string
	Competitors []RidersEntryRow
	TotalNations,
	TotalAthletes,
	TotalHorses int
}
