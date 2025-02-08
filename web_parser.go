package equestrian_events_riders_list

import (
	"github.com/gocolly/colly"
	"strings"
)

const (
	EquinisURL                     = "https://www.equinis.com/"
	TableClassSelector             = ".EventTable"
	TableRowsSelector              = "tr"
	RelativeEventDateClassSelector = "DateSectionTitleRow"
	EventRowClassSelector          = "EventRow"
)

// RelatedDateEventsTable group events by relative date. e.g. Upcoming, Recent
type RelatedDateEventsTable map[string][]EventEntryRow

// Parse retrieves the final list of riders and their horses.
func Parse() []*EquestrianCompetition {

	//call external API
	eventsTable, err := scrapEventsTable(EquinisURL)
	if err != nil {
		panic(err)
	}

	return GetEntryLists(eventsTable)
}

// scrapEventsTable retrieves the events from the table present in the main URL.
func scrapEventsTable(URL string) (events RelatedDateEventsTable, err error) {
	c := colly.NewCollector()
	events = make(RelatedDateEventsTable)

	c.OnHTML(TableClassSelector, func(e *colly.HTMLElement) {
		// traverse the table
		key := ""
		e.ForEach(TableRowsSelector, func(i int, e *colly.HTMLElement) {
			// if is title keep it as key
			if e.Attr("class") == RelativeEventDateClassSelector {
				key = e.Text
				events[e.Text] = []EventEntryRow{}
			}
			if e.Attr("class") == EventRowClassSelector {
				entryRow := EventEntryRow{}

				e.ForEach("td", func(idx int, elem *colly.HTMLElement) {
					//whole block checks if it is a link
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
				tmp := strings.TrimSpace(e.Text)
				evt := strings.Split(tmp, "\n\t")

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
		})
	})

	err = c.Visit(URL)
	return events, err
}
