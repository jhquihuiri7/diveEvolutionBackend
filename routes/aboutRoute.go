package routes

import (
	"diveEvolution/db"
	"diveEvolution/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)
var About *mongo.Collection
var AboutImg *mongo.Collection
func AboutHandler (w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	language := params["lang"]
	About = Client.Database("DiveEvolution").Collection("About")
	data := db.GetDocument(About,utils.GetLang(language, "about"))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func AboutImgHandler (w http.ResponseWriter, r *http.Request){
	AboutImg = Client.Database("DiveEvolution").Collection("AboutImg")
	data := db.GetDocument(AboutImg,"d04b58ce-52d4-4da4-a18b-9820689205ed")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
