package routes

import (
	"diveEvolution/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func ContactHandler(c *gin.Context) {
	language := c.Param("lang")
	f, err := os.Open("./data/DiveEvolution.Contact.json")
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(f)
	var contact []models.Contact
	err = decoder.Decode(&contact)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range contact {
		if v.GetLang(language) {
			c.Writer.WriteHeader(http.StatusOK)
			c.Writer.Header().Set("Content-Type", "application/json")
			d, _ := json.Marshal(v)
			c.Writer.Write(d)
			break
		}
	}
}
func ContactImgHandler(c *gin.Context) {
	f, err := os.Open("./data/DiveEvolution.ContactImg.json")
	if err != nil {
		log.Fatal(err)
	}

	var contactImg []models.ContactImg
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&contactImg)

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(contactImg[0])
	c.Writer.Write(d)
}
