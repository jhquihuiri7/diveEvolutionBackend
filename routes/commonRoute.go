package routes

import (
	"diveEvolution/db"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

var Footer *mongo.Collection
var HeaderImg *mongo.Collection
var FooterImg *mongo.Collection

func FooterHandler (w http.ResponseWriter, r *http.Request){
	Footer = Client.Database("DiveEvolution").Collection("Footer")
	data := db.GetDocument(Footer,"f93746d6-8b27-481f-ad1f-e888f7ef6d0f")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func HeaderImgHandler (w http.ResponseWriter, r *http.Request){
	HeaderImg = Client.Database("DiveEvolution").Collection("HeaderImg")
	data := db.GetDocument(HeaderImg,"7ec330b9-b663-4725-abb2-0786e28ca6ba")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func FooterImgHandler (w http.ResponseWriter, r *http.Request){
	FooterImg = Client.Database("DiveEvolution").Collection("FooterImg")
	data := db.GetDocument(FooterImg,"9756d287-94ae-47de-aacc-c454b9ff0ce8")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}