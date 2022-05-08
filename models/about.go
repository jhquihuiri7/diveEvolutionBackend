package models


type About struct {
	Id string `bson:"_id"`
	History `bson:"history"`
	Mission MVission `bson:"mission"`
	Vission MVission `bson:"vission"`
	Value Value `bson:"value"`
}
type History struct {
	Introduction string`bson:"introduction"`
	Brief string `bson:"brief"`
	Question string `bson:"question"`
	Button string `bson:"button"`
}
type MVission struct {
	Title string `bson:"title"`
	Description string `bson:"description"`
}
type Value struct {
	Title string `bson:"title"`
	Description string `bson:"description"`
	Values []string `bson:"values"`
}