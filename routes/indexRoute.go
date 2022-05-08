package routes

import (
	"diveEvolution/db"
	"diveEvolution/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)
var Index *mongo.Collection
var IndexImg *mongo.Collection

func IndexHandler (w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	language := params["lang"]
	Index = Client.Database("DiveEvolution").Collection("Index")
	data := db.GetDocument(Index,utils.GetLang(language,"index"))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func IndexImgHandler (w http.ResponseWriter, r *http.Request){
	IndexImg = Client.Database("DiveEvolution").Collection("IndexImg")
	data := db.GetDocument(IndexImg,"15388a1e-faaf-4fcb-9244-1ffa2813be59")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func UpdateIndexHandler (w http.ResponseWriter, r *http.Request)  {
	db.UpdateDocumment([]*mongo.Collection{Index, Footer})
}
