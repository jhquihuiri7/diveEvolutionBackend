package routes

import (
	"diveEvolution/db"
	"diveEvolution/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CoursesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	language := params["lang"]

	data := db.GetDocument(utils.Courses, utils.GetLang(language, "courses"))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func CoursesImgHandler(w http.ResponseWriter, r *http.Request) {
	data := db.GetDocument(utils.CoursesImg, "fd17c746-2dbb-4a51-9cd1-7becf48b2bf0")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
