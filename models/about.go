package models

import (
	"diveEvolution/utils"
)

type About struct {
	Id      string `json:"id"`
	History `json:"history"`
	Mission MVission `json:"mission"`
	Vission MVission `json:"vission"`
	Value   Value    `json:"value"`
}

func (a *About) GetLang(lang string) bool {
	if a.Id == utils.GetLang(lang, "about") {
		return true
	}
	return false
}

type History struct {
	Introduction string `json:"introduction"`
	Brief        string `json:"brief"`
	Question     string `json:"question"`
	Button       string `json:"button"`
	Ask          string `json:"ask"`
	Answer       string `json:"answer"`
}
type MVission struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
type Value struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Values      []string `json:"values"`
}

type AboutImg struct {
	Id         string `json:"id"`
	HistoryImg `json:"historyImg"`
	VissionImg string   `json:"vissionImg"`
	MissionImg string   `json:"missionImg"`
	ValueImg   []string `json:"valueImg"`
}
type HistoryImg struct {
	Background string `json:"background"`
	Logo       string `json:"logo"`
}
