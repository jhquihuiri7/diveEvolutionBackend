package routes

import (
	"diveEvolution/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func RunApp() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.GET("/api/getIndex/:lang", IndexHandler)
	router.GET("/api/getAbout/:lang", AboutHandler)
	router.GET("/api/getCourses/:lang", CoursesHandler)
	router.GET("/api/getContact/:lang", ContactHandler)
	router.GET("/api/getTours/:lang/:destination", ToursHandler)
	router.GET("/api/getFooter/es", FooterHandler)
	router.GET("/api/getCourseInfo/:lang/:ref", CoursesInfoHandler)
	//router.HandleFunc("/api/getToursInfo/{lang}/{ref}", ToursInfoHandler).Methods("GET")
	router.GET("/api/getIndexImg", IndexImgHandler)
	router.GET("/api/getAboutImg", AboutImgHandler)
	router.GET("/api/getCoursesImg", CoursesImgHandler)
	router.GET("/api/getContactImg", ContactImgHandler)
	router.GET("/api/getToursImg", ToursImgHandler)
	//router.GET("/api/getToursInfoImg/:ref", ToursInfoImgHandler)
	router.GET("/api/getHeaderImg", HeaderImgHandler)
	router.GET("/api/getFooterImg", FooterImgHandler)
	router.GET("/api/getCourseInfoImg/:ref", CoursesInfoImgHandler)

	//router.HandleFunc("/api/sendMail", MailSender).Methods("OPTIONS", "POST")

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, router)
}
