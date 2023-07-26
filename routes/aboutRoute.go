package routes

import (
	"diveEvolution/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func AboutHandler(c *gin.Context) {
	language := c.Param("lang")
	f, err := os.Open("./data/DiveEvolution.About.json")
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(f)
	var about []models.About
	err = decoder.Decode(&about)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range about {
		if v.GetLang(language) {
			c.Writer.WriteHeader(http.StatusOK)
			c.Writer.Header().Set("Content-Type", "application/json")
			d, _ := json.Marshal(v)
			c.Writer.Write(d)
			break
		}
	}
}
func AboutImgHandler(c *gin.Context) {
	f, err := os.Open("./data/DiveEvolution.AboutImg.json")
	if err != nil {
		log.Fatal(err)
	}

	var aboutImg []models.AboutImg
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&aboutImg)

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(aboutImg[0])
	c.Writer.Write(d)
}
