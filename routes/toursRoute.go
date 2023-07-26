package routes

import (
	"diveEvolution/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func ToursHandler(c *gin.Context) {
	language := c.Param("lang")
	destination := c.Param("destination")
	f, err := os.Open("./data/DiveEvolution.Tours.json")
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(f)
	var tours []models.Tours
	err = decoder.Decode(&tours)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range tours {
		if v.GetLang(language) {
			c.Writer.WriteHeader(http.StatusOK)
			c.Writer.Header().Set("Content-Type", "application/json")
			d, _ := json.Marshal(v.GetDestination(destination))
			c.Writer.Write(d)
			break
		}
	}
}
func ToursImgHandler(c *gin.Context) {
	f, err := os.Open("./data/DiveEvolution.ToursImg.json")
	if err != nil {
		log.Fatal(err)
	}

	var toursImg []models.ToursImg
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&toursImg)

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(toursImg[0])
	c.Writer.Write(d)
}
