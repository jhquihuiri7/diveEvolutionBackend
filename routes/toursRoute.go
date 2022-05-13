package routes

import (
	"diveEvolution/db"
	"diveEvolution/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)
var Tour *mongo.Collection
var TourImg *mongo.Collection
func ToursHandler (w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	language := params["lang"]
	destination := params["destination"]
	Tour = Client.Database("DiveEvolution").Collection("Tours")
	data := db.GetDocumentDest(Tour,utils.GetLang(language,"tours"),destination)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func ToursImgHandler (w http.ResponseWriter, r *http.Request){
	TourImg = Client.Database("DiveEvolution").Collection("ToursImg")
	data := db.GetDocument(TourImg,"3d34b3c3-94d8-45fb-bb36-3ca96c4a26c4")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
