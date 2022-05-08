package models

type Courses struct {
	Id string `bson:"_id"`
	Methodology
	CourseTypes []Course
}
type Methodology struct {
	Title string `bson:"title"`
	Methods []Method `bson:"methods"`
}
type Method struct {
	Title string `bson:"title"`
	Description string `bson:"description"`
	Question1 string `bson:"question1"`
	Works string `bson:"work"`
	Question2 string `bson:"question2"`
	Advantages []string `bson:"advantages"`
}
type Course struct {
	Ref string `bson:"ref"`
	Title string `bson:"title"`
	Description string `bson:"description"`
	Included []string `bson:"included"`
	Price int `bson:"price"`
}
type CoursesImg struct {
	Id string `bson:"_id"`
	MethodsImg []string `bson:"methodsImg"`
	CoursesImg []string  `bson:"coursesImg"`
}
