package routes

import (
	"diveEvolution/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func CoursesInfoHandler(c *gin.Context) {
	language := c.Param("lang")
	reference := c.Param("ref")

	f, err := os.Open("./data/DiveEvolution.CoursesInfo.json")
	if err != nil {
		log.Fatal(err)
	}
	var coursesInfo []models.CoursesInfo
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&coursesInfo)

	for _, v := range coursesInfo {
		if v.GetLang(language, "coursesInfo") {
			for _, cs := range v.CsInfo {
				if cs.GetCourse(reference) {
					fmt.Println(cs)
					c.Writer.WriteHeader(http.StatusOK)
					c.Writer.Header().Set("Content-Type", "application/json")
					d, _ := json.Marshal(cs)
					fmt.Println(string(d))
					c.Writer.Write(d)
					break
				}
			}
			break
		}
	}
}
func CoursesInfoImgHandler(c *gin.Context) {
	reference := c.Param("ref")

	f, err := os.Open("./data/DiveEvolution.CoursesInfoImg.json")
	if err != nil {
		log.Fatal(err)
	}

	var coursesInfoImg []models.CourseInfoImg
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&coursesInfoImg)

	for _, v := range coursesInfoImg {
		if v.GetCourse(reference) {
			c.Writer.WriteHeader(http.StatusOK)
			c.Writer.Header().Set("Content-Type", "application/json")
			d, _ := json.Marshal(v)
			c.Writer.Write(d)
		}
	}
}
