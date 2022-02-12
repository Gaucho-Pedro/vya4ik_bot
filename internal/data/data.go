package data

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

type CovidData struct {
	Date         string
	Hospitalized string
	Sick         string
	SickChange   string
	Healed       string
	HealedChange string
	Died         string
	DiedChange   string
}

func NewCovidData() *CovidData {
	return &CovidData{}
}

func (c *CovidData) GetData() {

	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered("https://стопкоронавирус.рф/information/", g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			c.Date = r.HTMLDoc.Find("div.cv-section__title-wrapper").Find("small").Last().Text()

			r.HTMLDoc.Find("div.cv-stats-virus__item").Each(func(i int, s *goquery.Selection) {
				switch i {
				case 0:
					c.Hospitalized = strings.Trim(s.Find("H3").Text(), " \n")
				case 1:
					c.HealedChange = strings.Trim(s.Find("H3").Text(), " \n")
				case 2:
					c.Healed = strings.Trim(s.Find("H3").Text(), " \n")
				case 3:
					c.SickChange = strings.Trim(s.Find("H3").Text(), " \n")
				case 4:
					c.Sick = strings.Trim(s.Find("H3").Text(), " \n")
				case 5:
					c.DiedChange = strings.Trim(s.Find("H3").Text(), " \n")
				case 6:
					c.Died = strings.Trim(s.Find("H3").Text(), " \n")
				}
			})
		},
		//BrowserEndpoint: "ws://localhost:3000",
	}).Start()
}
