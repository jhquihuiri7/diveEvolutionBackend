package models

type HeaderImg struct {
	Id   string `json:"id"`
	Logo string `json:"logo"`
}

type Footer struct {
	Id          string      `json:"id"`
	Phone       string      `json:"phone"`
	SocialMedia SocialMedia `json:"social_media"`
}

type SocialMedia struct {
	Facebook  string `json:"facebook"`
	Instagram string `json:"instagram"`
	Twitter   string `json:"twitter"`
}

type FooterImg struct {
	Id               string           `json:"id"`
	Logo             string           `json:"logo"`
	Background       string           `json:"background"`
	Whatsapp         string           `json:"whatsapp"`
	SocialMediaIcons SocialMediaIcons `json:"social_media_icons"`
}

type SocialMediaIcons struct {
	Facebook  string `json:"facebook"`
	Instagram string `json:"instagram"`
	Twitter   string `json:"twitter"`
}
