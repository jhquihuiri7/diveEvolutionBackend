package models

import "diveEvolution/utils"

type CoursesInfo struct {
	Id     string       `json:"id"`
	CsInfo []CourseInfo `json:"coursesInfo"`
}

func (ci *CoursesInfo) GetLang(lang, route string) bool {
	if ci.Id == utils.GetLang(lang, route) {
		return true
	}
	return false
}

type CourseInfo struct {
	Course           `json:"course"`
	LargeDescription string   `json:"largeDescription"`
	Itinerary        []string `json:"itinerary"`
}

func (ci *CourseInfo) GetCourse(ref string) bool {
	if ci.Course.Ref == ref {
		return true
	}
	return false
}

type CourseInfoImg struct {
	Ref string   `json:"ref"`
	Img []string `json:"img"`
}

func (ci *CourseInfoImg) GetCourse(ref string) bool {
	if ci.Ref == ref {
		return true
	}
	return false
}
