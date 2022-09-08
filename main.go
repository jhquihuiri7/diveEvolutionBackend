package main

import (
	"diveEvolution/db"
	"diveEvolution/routes"
	"diveEvolution/utils"
)

func init() {
	utils.Client = db.ConnectDB()
	utils.Index = utils.Client.Database("DiveEvolution").Collection("Index")
	utils.IndexImg = utils.Client.Database("DiveEvolution").Collection("IndexImg")
	utils.About = utils.Client.Database("DiveEvolution").Collection("About")
	utils.AboutImg = utils.Client.Database("DiveEvolution").Collection("AboutImg")
	utils.Footer = utils.Client.Database("DiveEvolution").Collection("Footer")
	utils.HeaderImg = utils.Client.Database("DiveEvolution").Collection("HeaderImg")
	utils.FooterImg = utils.Client.Database("DiveEvolution").Collection("FooterImg")
	utils.Contact = utils.Client.Database("DiveEvolution").Collection("Contact")
	utils.ContactImg = utils.Client.Database("DiveEvolution").Collection("ContactImg")
	utils.CoursesInfo = utils.Client.Database("DiveEvolution").Collection("CoursesInfo")
	utils.CoursesInfoImg = utils.Client.Database("DiveEvolution").Collection("CoursesInfoImg")
	utils.Courses = utils.Client.Database("DiveEvolution").Collection("Courses")
	utils.CoursesImg = utils.Client.Database("DiveEvolution").Collection("CoursesImg")
	utils.ToursInfo = utils.Client.Database("DiveEvolution").Collection("ToursInfo")
	utils.ToursInfoImg = utils.Client.Database("DiveEvolution").Collection("ToursInfoImg")
	utils.Tour = utils.Client.Database("DiveEvolution").Collection("Tours")
	utils.TourImg = utils.Client.Database("DiveEvolution").Collection("ToursImg")
}
func main() {
	routes.RunApp()
}
