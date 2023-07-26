package routes

import (
	"diveEvolution/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func HeaderImgHandler(c *gin.Context) {
	f, err := os.Open("./data/DiveEvolution.HeaderImg.json")
	if err != nil {
		log.Fatal(err)
	}

	var indexImg []models.HeaderImg
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&indexImg)

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(indexImg[0])
	c.Writer.Write(d)
}
