package routes

import (
	"diveEvolution/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func CoursesHandler(c *gin.Context) {
	language := c.Param("lang")
	f, err := os.Open("./data/DiveEvolution.Courses.json")
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(f)
	var courses []models.Courses
	err = decoder.Decode(&courses)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range courses {
		if v.GetLang(language) {
			c.Writer.WriteHeader(http.StatusOK)
			c.Writer.Header().Set("Content-Type", "application/json")
			d, _ := json.Marshal(v)
			c.Writer.Write(d)
			break
		}
	}
}
func CoursesImgHandler(c *gin.Context) {
	f, err := os.Open("./data/DiveEvolution.CoursesImg.json")
	if err != nil {
		log.Fatal(err)
	}

	var coursesImg []models.CoursesImg
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&coursesImg)

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(coursesImg[0])
	c.Writer.Write(d)
}
