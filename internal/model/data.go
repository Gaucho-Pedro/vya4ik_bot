package model

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
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

	response, err := http.Get("https://стопкоронавирус.рф/information/")
	if err != nil {
		log.Panic(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panic("Error reading HTTP body: ", err)
	}
	log.Info("Start parsing...")
	c.Date = strings.Trim(regexp.MustCompile(">П(.)*?\\d</s").FindString(string(body)), "></s")

	res := strings.NewReplacer("+", "", "'", "", "=", "").Replace(regexp.MustCompile("='(.)*?'").FindString(string(body)))
	if err := json.Unmarshal([]byte(res), c); err != nil {
		log.Panic(err)
	}
	log.Info("Parsing successful!")
}
