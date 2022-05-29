package models

type CourseInfo struct {
	Id string `bson:"_id"`
	Ref string `bson:"ref"`
	Lang string `bson:"lang"`
	Course 	`bson:"course"`
	LargeDescription string `bson:"largeDescription"`
	Itinerary []string `bson:"itinerary"`
}
type CourseInfoImg struct {
	Id string `bson:"_id"`
	Img []string `bson:"img"`
}
