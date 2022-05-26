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

var CoursesInfo *mongo.Collection
var CoursesInfoImg *mongo.Collection
func CoursesInfoHandler (w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	language := params["lang"]
	reference := params["ref"]
	CoursesInfo = Client.Database("DiveEvolution").Collection("CoursesInfo")
	data := db.GetDocumentCourse(CoursesInfo, language, reference)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func CoursesInfoImgHandler (w http.ResponseWriter, r *http.Request){
	CoursesInfoImg = Client.Database("DiveEvolution").Collection("CoursesImg")
	data := db.GetDocument(CoursesInfoImg,"fd17c746-2dbb-4a51-9cd1-7becf48b2bf0")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}

