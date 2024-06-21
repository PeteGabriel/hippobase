package equestrian_events_riders_list

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

func GetEntryLists(events RelatedDateEventsTable) []EventInfo {
	var parsedEvents []EventInfo

	for _, v := range events {
		for _, event := range v {

			if event.EntryListURL == "" {
				continue
			}

			e, err := parseEvent(event.EntryListURL)
			if err != nil {
				panic(err)
			}
			parsedEvents = append(parsedEvents, e)
		}
	}

	return parsedEvents
}

func parseEvent(eventURL string) (EventInfo, error) {
	eventInfo := EventInfo{}

	c := colly.NewCollector()

	//event name
	c.OnHTML(".EventTitle", func(e *colly.HTMLElement) {
		eventInfo.EventFullName = e.Text
	})
	//creation date of the list
	c.OnHTML(".CreationDate", func(e *colly.HTMLElement) {
		eventInfo.CreatedAt = e.Text
	})
	//get countries info
	c.OnHTML(".CountryTable", func(e *colly.HTMLElement) {
		entryList, _ := parseEntryList(e)
		eventInfo.Competitors = entryList
	})
	//get numbers summary
	c.OnHTML(".NumberSummary", func(e *colly.HTMLElement) {
		totalNations, totalAthletes, totalHorses := parseNumbersSummary(e)
		eventInfo.TotalNations = totalNations
		eventInfo.TotalAthletes = totalAthletes
		eventInfo.TotalHorses = totalHorses
	})

	err := c.Visit(eventURL)
	return eventInfo, err
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
