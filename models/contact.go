package models

import "diveEvolution/utils"

type Contact struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
	ContactInfo  `json:"contactInfo"`
	Button       string `json:"button"`
	FormTitle    string `json:"formTitle"`
}

func (c *Contact) GetLang(lang string) bool {
	if c.Id == utils.GetLang(lang, "contact") {
		return true
	}
	return false
}

type ContactInfo struct {
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}
type ContactImg struct {
	Id         string `json:"id"`
	Background string `json:"background"`
	Foreground string `json:"foreground"`
	Form       string `json:"form"`
}
