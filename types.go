package equestrian_events_riders_list

// RidersEntryRow represents a row in the entry list of riders for a certain event.
type RidersEntryRow struct {
	Flag        string `json:"flag"`
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	// horse rider - set of horses
	Pairs map[string][]string `json:"pairs"`
}

// EventEntryRow represents a row in the minified event table.
type EventEntryRow struct {
	Date, Name, Location, EventURL, EntryListURL string
}

// EventInfo represents the complete information of an event.
// It contains the event name, the date of creation and the entry list of competitors.
type EventInfo struct {
	CreatedAt     string           `json:"created_at"`
	EventFullName string           `json:"event_name"`
	Competitors   []RidersEntryRow `json:"competitors,omitempty"`
	TotalNations  int              `json:"total_nations"`
	TotalAthletes int              `json:"total_athletes"`
	TotalHorses   int              `json:"total_horses"`
}
