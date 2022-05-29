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
	params:= mux.Vars(r)
	reference := params["ref"]
	CoursesInfoImg = Client.Database("DiveEvolution").Collection("CoursesInfoImg")
	data := db.GetDocument(CoursesInfoImg,reference)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}

