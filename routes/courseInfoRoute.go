package routes

import (
	"diveEvolution/db"
	"diveEvolution/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CoursesInfoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	language := params["lang"]
	reference := params["ref"]
	data := db.GetDocumentCourse(utils.CoursesInfo, language, reference)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func CoursesInfoImgHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	reference := params["ref"]
	data := db.GetDocument(utils.CoursesInfoImg, reference)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
