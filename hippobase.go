package equestrian_events_riders_list

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

func GetEntryLists(events RelatedDateEventsTable) []*EquestrianCompetition {
	var competitions []*EquestrianCompetition

	for _, v := range events {
		equestrianCompetition := &EquestrianCompetition{}
		parsedEvents := make([]*EventInfo, 0)
		for _, event := range v {

			if event.EntryListURL == "" {
				continue
			}

			e, err := parseCompetition(equestrianCompetition, event.EntryListURL)
			if err != nil {
				panic(err)
			}
			for _, evt := range e {
				parsedEvents = append(parsedEvents, evt)
			}

		}
		equestrianCompetition.events = parsedEvents
		competitions = append(competitions, equestrianCompetition)
	}

	return competitions
}

// for a specific competition, returns a series of events held during that.
func parseCompetition(comp *EquestrianCompetition, eventURL string) ([]*EventInfo, error) {
	events := make([]*EventInfo, 0)
	c := colly.NewCollector()

	//get title for the entire competition
	c.OnHTML(".EventTitle", func(e *colly.HTMLElement) {
		comp.MainTitle = strings.TrimSpace(e.Text)
	})

	//scrap the different parts of the competition
	c.OnHTML(".Content", func(e *colly.HTMLElement) {

		e.ForEach(".EntryGroupTitle", func(i int, e *colly.HTMLElement) {
			var eventInfo *EventInfo
			if len(events)-1 >= i {
				eventInfo = events[i]
			} else {
				eventInfo = &EventInfo{}
			}

			//format name
			eventInfo.EventFullName = strings.TrimSpace(e.Text)
			eventInfo.EventFullName = strings.Replace(eventInfo.EventFullName, "\n\t\t\t", " ", -1)

			events = append(events, eventInfo)
		})

		//inside content there are different blocks that represent different events held during the competition
		e.ForEach(".EntryGroup", func(i int, e *colly.HTMLElement) {
			eventInfo := events[i]
			entryList, _ := parseEntryList(e)
			eventInfo.Competitors = entryList
		})

		//get the total number of nations, athletes and horses
		e.ForEach(".NumberSummary", func(i int, e *colly.HTMLElement) {
			eventInfo := events[i]
			tNations, athletes, horses := parseNumbersSummary(e)
			eventInfo.TotalNations = tNations
			eventInfo.TotalAthletes = athletes
			eventInfo.TotalHorses = horses
		})

		comp.events = events
	})

	err := c.Visit(eventURL)
	return events, err
}

func parseEntryList(e *colly.HTMLElement) ([]RidersEntryRow, error) {
	entryList := make([]RidersEntryRow, 0)

	e.ForEach(".CountryBlock", func(i int, row *colly.HTMLElement) {
		riderHorse := RidersEntryRow{}
		row.DOM.Find(".CountryRow").Each(func(i int, selection *goquery.Selection) {
			parseCountryRow(selection, &riderHorse)
		})

		row.DOM.Find(".CompetitorRow").Each(func(i int, selection *goquery.Selection) {
			if selection.Nodes[0].Attr[0].Val == "CompetitorRow SeparatorRow" { // ignore separation row
				return
			}
			parseCompetitorRow(selection, &riderHorse)
			entryList = append(entryList, riderHorse)
		})
	})

	return entryList, nil
}

func parseNumbersSummary(e *colly.HTMLElement) (int, int, int) {
	t := strings.TrimSpace(e.Text)
	t = strings.Replace(t, "\n\t\t", ":", -1)
	t = strings.Replace(t, " ", "", -1)
	parts := strings.Split(t, ":")

	tNations, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	athletes, err := strconv.Atoi(parts[3])
	if err != nil {
		panic(err)
	}
	horses, err := strconv.Atoi(parts[5])
	if err != nil {
		panic(err)
	}

	return tNations, athletes, horses
}

func parseCountryRow(countryRow *goquery.Selection, entry *RidersEntryRow) {
	countryRow.Find("td").Each(func(i int, innerSelection *goquery.Selection) {
		isV, _ := innerSelection.Attr("class")
		if isV == "CountrySymbols" {
			entry.CountryCode = strings.TrimSpace(innerSelection.Text())
			entry.Flag, _ = innerSelection.Find("img").First().Attr("src")
			return
		}
		if isV == "CountryName" {
			entry.CountryName = innerSelection.Text()
			return
		}
	})
}

func parseCompetitorRow(competitorRow *goquery.Selection, entry *RidersEntryRow) {
	pairs := make(map[string][]string, 0)
	riderName := ""
	competitorRow.Find("td").Each(func(i int, innerSelection *goquery.Selection) {
		isV, _ := innerSelection.Attr("class")
		if isV == "Competitor" {
			riderName = innerSelection.Text()
			pairs[riderName] = []string{}
			return
		}
		if isV == "HorseCell" {
			innerSelection.Find(".HName").Each(func(i int, selection *goquery.Selection) {
				pairs[riderName] = append(pairs[riderName], selection.Text())
			})
			return
		}
	})
	entry.Pairs = pairs
}
