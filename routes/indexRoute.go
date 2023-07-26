package routes

import (
	"diveEvolution/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func IndexHandler(c *gin.Context) {
	language := c.Param("lang")
	f, err := os.Open("./data/DiveEvolution.Index.json")
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(f)
	var index []models.Index
	err = decoder.Decode(&index)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range index {
		if v.GetLang(language) {
			c.Writer.WriteHeader(http.StatusOK)
			c.Writer.Header().Set("Content-Type", "application/json")
			d, _ := json.Marshal(v)
			c.Writer.Write(d)
			break
		}
	}
}

func IndexImgHandler(c *gin.Context) {
	f, err := os.Open("./data/DiveEvolution.IndexImg.json")
	if err != nil {
		log.Fatal(err)
	}

	var indexImg []models.IndexImg
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&indexImg)

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(indexImg[0])
	c.Writer.Write(d)
}
