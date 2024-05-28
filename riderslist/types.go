package riderslist

// RidersEntryRow represents a row in the entry list of riders for a certain event.
type RidersEntryRow struct {
	flag, countryCode, countryName string
	// horse rider - set of horses
	pairs map[string][]string
}

// EventEntryRow represents a row in the minified event table.
type EventEntryRow struct {
	date, name, location, eventURL, entryListURL string
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
