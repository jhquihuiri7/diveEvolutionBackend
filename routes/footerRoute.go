package routes

import (
	"diveEvolution/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func FooterHandler(c *gin.Context) {
	f, err := os.Open("./data/DiveEvolution.Footer.json")
	if err != nil {
		log.Fatal(err)
	}

	var footer []models.Footer
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&footer)

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(footer[0])
	c.Writer.Write(d)
}

func FooterImgHandler(c *gin.Context) {
	f, err := os.Open("./data/DiveEvolution.FooterImg.json")
	if err != nil {
		log.Fatal(err)
	}

	var footerImg []models.FooterImg
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&footerImg)

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(footerImg[0])
	c.Writer.Write(d)
}
