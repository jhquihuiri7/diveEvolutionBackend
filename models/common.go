package models

type Header struct {
	Id     string `bson:"_id"`
	NavBar `bson:"navBar"`
	Title  string `bson:"title"`
	Button string `bson:"button"`
}
type Footer struct {
	Id          string `bson:"_id"`
	NavBar      `bson:"navBar"`
	Phone       string `bson:"phone"`
	SocialMedia `bson:"social_media"`
	Copyright   string `bson:"copyright"`
}
type SocialMedia struct {
	Facebook  string `bson:"facebook"`
	Instagram string `bson:"instagram"`
	Twitter   string `bson:"twitter"`
}
type NavBar struct {
	NavBarItems []NavBarItem `bson:"nav_bar_items"`
}
type NavBarItem struct {
	Title     string   `bson:"title"`
	Accordion []string `bson:"accordion"`
}

///IMG///
type HeaderImg struct {
	Id        string `bson:"_id"`
	Logo      string `bson:"logo"`
	LangItems []LangItem
}
type FooterImg struct {
	Id               string `bson:"_id"`
	Logo             string `bson:"logo"`
	Background       string `bson:"background"`
	WhatsApp         string `bson:"whatsapp"`
	SocialMediaIcons `bson:"social_media_icons"`
}
type LangItem struct {
	Img  string `bson:"img"`
	Lang string `bson:"lang"`
}
type SocialMediaIcons struct {
	Facebook  string `bson:"facebook"`
	Instagram string `bson:"instagram"`
	Twitter   string `bson:"twitter"`
}

///https://via.placeholder.com/1050x200/0000FF/808080%20?Text=Digital.com
