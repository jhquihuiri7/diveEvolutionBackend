package models

import "diveEvolution/utils"

type Tours struct {
	Id string   `json:"id"`
	Sc TourType `json:"sc"`
	Sx TourType `json:"sx"`
	Ib TourType `json:"ib"`
}

func (c *Tours) GetLang(lang string) bool {
	if c.Id == utils.GetLang(lang, "tours") {
		return true
	}
	return false
}
func (c *Tours) GetDestination(destination string) interface{} {
	if destination == "sc" {
		return c.Sc
	} else if destination == "sx" {
		return c.Sx
	} else if destination == "ib" {
		return c.Ib
	}
	return c.Sc
}

type TourType struct {
	Snorkel []Course `json:"snorkel"`
	Dive    []Course `json:"dive"`
	Land    []Course `json:"land"`
}
type ToursImg struct {
	Id         string   `json:"id"`
	SnorkelImg []string `json:"snorkelImg"`
	DiveImg    []string `json:"diveImg"`
	LandImg    []string `json:"landImg"`
}
