package routes

import (
	"diveEvolution/db"
	"github.com/gorilla/mux"
	//"diveEvolution/utils"
	"encoding/json"
	//"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

var ToursInfo *mongo.Collection
var ToursInfoImg *mongo.Collection
func ToursInfoHandler (w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	language := params["lang"]
	reference := params["ref"]
	ToursInfo = Client.Database("DiveEvolution").Collection("ToursInfo")
	data := db.GetDocumentCourse(ToursInfo, language, reference)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func ToursInfoImgHandler (w http.ResponseWriter, r *http.Request){
	params:= mux.Vars(r)
	reference := params["ref"]
	ToursInfoImg = Client.Database("DiveEvolution").Collection("ToursInfoImg")
	data := db.GetDocument(ToursInfoImg,reference)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}

