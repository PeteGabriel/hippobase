package hippobase

import (
	"github.com/gocolly/colly"
	"strings"
)

const (
	// EquinisURL is the main URL to scrap the events table.
	EquinisURL = "https://www.equinis.com/"
	// TableClassSelector is the class selector for the events table.
	TableClassSelector = ".EventTable"
	// TableRowsSelector is the selector for the rows of the table.
	TableRowsSelector = "tr"
	// RelativeEventDateClassSelector is the class selector for the relative event date.
	RelativeEventDateClassSelector = "DateSectionTitleRow"
	// EventRowClassSelector is the class selector for the event row.
	EventRowClassSelector = "EventRow"
)

// RelatedDateEventsTable group events by relative date. e.g. Upcoming, Recent
type RelatedDateEventsTable map[string][]EventEntryRow

// Parse retrieves the final list of riders and their horses.
func Parse() []*EquestrianCompetition {

	//call external API
	eventsTable, err := scrapTable(EquinisURL)
	if err != nil {
		panic(err)
	}

	return GetEntryLists(eventsTable)
}

// scrapTable retrieves the events from the table present in the main URL.
func scrapTable(URL string) (events RelatedDateEventsTable, err error) {
	c := colly.NewCollector()
	events = make(RelatedDateEventsTable)

	c.OnHTML(TableClassSelector, scrapEvents(events))

	err = c.Visit(URL)
	return events, err
}

func scrapEvents(events RelatedDateEventsTable) func(e *colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		// traverse the table
		key := ""
		e.ForEach(TableRowsSelector, scrapEachRow(events, key))
	}
}

func scrapEachRow(events RelatedDateEventsTable, key string) func(i int, e *colly.HTMLElement) {
	return func(i int, e *colly.HTMLElement) {
		// if is a title, keep it as key
		if e.Attr("class") == RelativeEventDateClassSelector {
			key = e.Text
			events[e.Text] = []EventEntryRow{}
		}
		if e.Attr("class") == EventRowClassSelector {
			entryRow := EventEntryRow{}

			e.ForEach("td", func(idx int, elem *colly.HTMLElement) {
				//this whole block checks if it is a link
				if elem.Text == "" {
					if len(elem.DOM.Nodes) > 0 {
						if elem.DOM.Nodes[0].FirstChild != nil && elem.DOM.Nodes[0].FirstChild.Data == "a" {
							// check if it is event homepage URL
							if elem.DOM.Nodes[0].FirstChild.Attr[1].Val == "EventPage" {
								// both contain the same html properties, but they always follow the same order.
								// just assign content to the second if the first is not empty

								if entryRow.EventURL == "" {
									entryRow.EventURL = elem.DOM.Nodes[0].FirstChild.Attr[0].Val

									//force https
									if !strings.HasPrefix(entryRow.EventURL, "https") {
										entryRow.EventURL = strings.Replace(entryRow.EventURL, "http", "https", 1)
									}
								} else if entryRow.EntryListURL == "" {
									entryRow.EntryListURL = elem.DOM.Nodes[0].FirstChild.Attr[0].Val

									//force https
									if !strings.HasPrefix(entryRow.EntryListURL, "https") {
										entryRow.EntryListURL = strings.Replace(entryRow.EntryListURL, "http", "https", 1)
									}
								}
							}
						}
					}
				}
			})

			// info contains a lot of spaces and new lines
			trimmed := strings.TrimSpace(e.Text)
			evt := strings.Split(trimmed, "\n\t")

			switch len(evt) {
			case 1:
				entryRow.Date = strings.TrimSpace(evt[0])
			case 2:
				entryRow.Date = strings.TrimSpace(evt[0])
				entryRow.Name = strings.TrimSpace(evt[1])
			case 3:
				entryRow.Date = strings.TrimSpace(evt[0])
				entryRow.Name = strings.TrimSpace(evt[1])
				entryRow.Location = strings.TrimSpace(evt[2])
			}
			events[key] = append(events[key], entryRow)
		}
	}
}
