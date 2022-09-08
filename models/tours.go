package models

type Tours struct {
	Id string   `bson:"_id"`
	Sc TourType `bson:"sc"`
	Sx TourType `bson:"sx"`
	Ib TourType `bson:"ib"`
}
type TourType struct {
	Snorkel []Course `bson:"snorkel"`
	Dive    []Course `bson:"dive"`
	Land    []Course `bson:"land"`
}
type ToursImg struct {
	Id         string   `bson:"_id"`
	SnorkelImg []string `bson:"snorkelImg"`
	DiveImg    []string `bson:"diveImg"`
	LandImg    []string `bson:"landImg"`
}
