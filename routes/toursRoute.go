package routes

import (
	"diveEvolution/db"
	"diveEvolution/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func ToursHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	language := params["lang"]
	destination := params["destination"]
	data := db.GetDocumentDest(utils.Tour, utils.GetLang(language, "tours"), destination)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func ToursImgHandler(w http.ResponseWriter, r *http.Request) {
	data := db.GetDocument(utils.TourImg, "3d34b3c3-94d8-45fb-bb36-3ca96c4a26c4")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
