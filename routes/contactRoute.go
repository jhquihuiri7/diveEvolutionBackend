package routes

import (
	"diveEvolution/db"
	"diveEvolution/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)
 var Contact *mongo.Collection
var ContactImg *mongo.Collection
func ContactHandler (w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	language := params["lang"]
	Contact = Client.Database("DiveEvolution").Collection("Contact")
	data := db.GetDocument(Contact,utils.GetLang(language,"contact"))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func ContactImgHandler (w http.ResponseWriter, r *http.Request){
	ContactImg = Client.Database("DiveEvolution").Collection("ContactImg")
	data := db.GetDocument(ContactImg,"3aea6d6c-3fac-4b11-8ee0-5f6a6839da74")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
