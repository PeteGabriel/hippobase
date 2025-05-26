package internal

import (
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

func ParseEvent(entryListURL string) (*EquestrianCompetition, error) {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	competition := &EquestrianCompetition{
		MainTitle: "",
		Events:    make([]*EventInfo, 0),
	}

	c.OnHTML(".EventTitle", func(e *colly.HTMLElement) {
		competition.MainTitle = strings.TrimSpace(e.Text)
	})

	c.OnHTML(".EntryGroup", parseSingleEntryGroup(competition))

	err := c.Visit(entryListURL)
	if err != nil {
		return nil, err
	}
	return competition, nil
}

func parseSingleEntryGroup(competition *EquestrianCompetition) func(e *colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		// Parse the entry group title
		entryGroupTitle := strings.TrimSpace(e.ChildText(".EntryGroupTitle"))
		entryGroupTitle = strings.Replace(entryGroupTitle, "\n\t\t\t", " ", -1)

		creationDate := strings.TrimSpace(e.ChildText(".CreationDate"))
		creationDate = strings.Replace(creationDate, "\n\t\t\t", " ", -1)

		entry := &EventInfo{
			EventFullName: entryGroupTitle,
			CreatedAt:     creationDate,
			Competitors:   make([]*RidersEntryRow, 0),
		}

		// Parse the entries within the group
		e.ForEach(".CountryBlock", func(i int, row *colly.HTMLElement) {
			// Parse country info
			countryName := strings.TrimSpace(row.ChildText(".CountryName"))
			countryName = strings.Replace(countryName, "\n\t\t\t", " ", -1)
			countryID := strings.TrimSpace(row.ChildText(".CountryID"))
			countryID = strings.Replace(countryID, "\n\t\t\t", " ", -1)
			flag := strings.TrimSpace(row.ChildAttr("img", "src"))

			riderRow := &RidersEntryRow{
				Flag:        flag,
				CountryCode: countryID,
				CountryName: countryName,
				Pairs:       make([]*CompetitorHorsePair, 0),
			}

			row.ForEach(".CompetitorRow", func(i int, competitor *colly.HTMLElement) {
				name := strings.TrimSpace(competitor.ChildText(".Competitor"))
				// dummy validation because of how web is structured
				if name == "" {
					return
				}

				hs := make([]string, 0)
				competitor.ForEach(".Horse", func(i int, horse *colly.HTMLElement) {
					//remove special characters
					h := strings.ReplaceAll(strings.TrimSpace(horse.Text), ",", "")
					hs = append(hs, strings.TrimSpace(h))
				})

				riderRow.Pairs = append(riderRow.Pairs, &CompetitorHorsePair{
					Competitor: name,
					Horses:     hs,
				})
			})

			entry.Competitors = append(entry.Competitors, riderRow)
		})

		e.ForEach(".NumberSummary", func(i int, elem *colly.HTMLElement) {
			if i == 0 {
				entry.TotalNations, entry.TotalAthletes, entry.TotalHorses = parseNumbersSummary(elem)
			}
		})

		competition.Events = append(competition.Events, entry)
	}
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
