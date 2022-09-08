package models

type Index struct {
	Id   string `bson:"_id"`
	Body `bson:"body"`
}
type Body struct {
	Section1 `bson:"section1"`
	Section2 `bson:"section2"`
	Title    string `bson:"title"`
	Button   string `bson:"button"`
}

type Section1 struct {
	Calidad []string `bson:"calidad"`
	Precio  []string `bson:"precio"`
}
type Section2 struct {
	Items []ItemsInfo `bson:"items"`
}
type ItemsInfo struct {
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Button      string `bson:"button"`
}

type IndexImg struct {
	Id      string `bson:"_id"`
	BodyImg `bson:"body_img"`
}
type BodyImg struct {
	Section1    `bson:"section1"`
	Section2Img `bson:"section2"`
	Background  string `bson:"backgroud"`
}
type Section2Img struct {
	ItemsImg []string `bson:"items_img"`
}
