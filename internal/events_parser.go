package internal

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

// equinis is the main URL to scrap the events table.
const equinis = "https://www.equinis.com/"

type Events []EventEntryRow

// ParseMainEvents ParseMainTable retrieves the final list of events from the main table.
// It returns a slice of EventEntryRow, which contains the event details.
// The function is designed to be called from the main function or any other part of the code where event data is needed.
func ParseMainEvents() (Events, error) {
	return scrapMainTable(equinis)
}

// FirstEntryListURL returns the first entry list URL from the events.
// Sometimes the first event does not have an entry list URL, so it returns the first available one.
func (e Events) FirstEntryListURL() string {
	if len(e) == 0 {
		return ""
	}
	for _, entry := range e {
		if entry.EntryListURL != "" {
			return entry.EntryListURL
		}
	}
	return ""
}

func scrapMainTable(url string) (Events, error) {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	events := make(Events, 0)

	c.OnHTML(`.EventRow`, func(e *colly.HTMLElement) {
		event := EventEntryRow{}
		sanitizeName(e, &event)

		// add event URL and entry list URL
		targets := e.DOM.Find(`a[target]`)
		targets.Each(func(i int, s *goquery.Selection) {
			val, exists := s.Attr("target")
			if exists && val == "EventPage" {
				if event.EventURL == "" {
					event.EventURL, _ = s.Attr("href")
				} else {
					event.EntryListURL, _ = s.Attr("href")
					// grab ID from the entry list URL
					event.Id = getIdFromEntryList(event.EntryListURL)
				}
			}
		})

		events = append(events, event)
	})

	err := c.Visit(url)

	return events, err
}

func getIdFromEntryList(url string) int {
	p := strings.Split(url, "EventID=")
	id, _ := strconv.Atoi(p[1])
	return id
}

// names contain a lot of empty spaces and new lines.
func sanitizeName(e *colly.HTMLElement, event *EventEntryRow) {
	trimmed := strings.TrimSpace(e.Text)
	evt := strings.Split(trimmed, "\n\t")

	switch len(evt) {
	case 1:
		event.Date = strings.TrimSpace(evt[0])
	case 2:
		event.Date = strings.TrimSpace(evt[0])
		event.Name = strings.TrimSpace(evt[1])
	case 3:
		event.Date = strings.TrimSpace(evt[0])
		event.Name = strings.TrimSpace(evt[1])
		event.Location = strings.TrimSpace(evt[2])
	}
}
