package hippobase

import "github.com/petegabriel/hippobase/internal"

// GetEvents retrieves the events from the Hippobase website.
// It uses the colly library to scrape the data from the specified URL.
func GetEvents() (internal.Events, error) {
	events, err := internal.ParseMainEvents()
	if err != nil {
		return nil, err
	}
	return events, nil
}

// GetEntryLists retrieves the entry lists for a specific event URL.
func GetEntryLists(eventURL string) (*internal.EquestrianCompetition, error) {
	return internal.ParseEvent(eventURL)
}
