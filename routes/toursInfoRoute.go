package routes

import (
	"diveEvolution/db"
	"diveEvolution/utils"
	"github.com/gorilla/mux"
	//"diveEvolution/utils"
	"encoding/json"

	"net/http"
)

func ToursInfoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	language := params["lang"]
	reference := params["ref"]
	data := db.GetDocumentCourse(utils.ToursInfo, language, reference)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func ToursInfoImgHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	reference := params["ref"]
	data := db.GetDocument(utils.ToursInfoImg, reference)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
