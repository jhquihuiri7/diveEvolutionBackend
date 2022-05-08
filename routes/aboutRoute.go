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
