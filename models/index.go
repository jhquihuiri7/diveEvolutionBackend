package models

import (
	"diveEvolution/utils"
	"fmt"
)

type Index struct {
	Id   string `json:"id"`
	Body `json:"body"`
}

func (i *Index) GetLang(lang string) bool {
	fmt.Println(i.Id)
	if i.Id == utils.GetLang(lang, "index") {
		return true
	}
	return false
}

type Body struct {
	Section1 `json:"section1"`
	Section2 `json:"section2"`
	Title    string `json:"title"`
	Button   string `json:"button"`
}

type Section1 struct {
	Calidad []string `json:"calidad"`
	Precio  []string `json:"precio"`
}
type Section2 struct {
	Items []*ItemsInfo `json:"items"`
}
type ItemsInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Button      string `json:"button"`
}

type IndexImg struct {
	BodyImg `json:"body_img"`
}
type BodyImg struct {
	Section1    `json:"section1"`
	Section2Img `json:"section2"`
	Background  string `json:"backgroud"`
}
type Section2Img struct {
	ItemsImg []string `json:"items"`
}
