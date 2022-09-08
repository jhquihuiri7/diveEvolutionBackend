package models

type Contact struct {
	Id           string `bson:"_id"`
	Title        string `bson:"title"`
	Introduction string `bson:"introduction"`
	ContactInfo  `bson:"contactInfo"`
	Button       string `bson:"button"`
	FormTitle    string `bs:"formTitle"`
}
type ContactInfo struct {
	Address string `bson:"address"`
	Email   string `bson:"email"`
	Phone   string `bson:"phone"`
}
type ContactImg struct {
	Id         string `bson:"_id"`
	Background string `bson:"background"`
	Foreground string `bson:"foreground"`
	Form       string `bson:"form"`
}
