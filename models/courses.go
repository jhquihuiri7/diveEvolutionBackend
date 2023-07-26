package models

import "diveEvolution/utils"

type Courses struct {
	Id          string `json:"id"`
	Methodology `json:"methodology"`
	CourseTypes []Course `json:"courseTypes"`
}

func (c *Courses) GetLang(lang string) bool {
	if c.Id == utils.GetLang(lang, "courses") {
		return true
	}
	return false
}

type Methodology struct {
	Title   string   `json:"title"`
	Methods []Method `json:"methods"`
}
type Method struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Question1   string   `json:"question1"`
	Works       string   `json:"work"`
	Question2   string   `json:"question2"`
	Advantages  []string `json:"advantages"`
}
type Course struct {
	Ref         string   `json:"ref"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Included    []string `json:"included"`
	Price       int      `json:"price"`
}
type CoursesImg struct {
	Id         string   `json:"id"`
	MethodsImg []string `json:"methodsImg"`
	CoursesImg []string `json:"coursesImg"`
}
