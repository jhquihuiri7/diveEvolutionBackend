package routes

import (
	"diveEvolution/db"
	"diveEvolution/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	language := params["lang"]
	data := db.GetDocument(utils.About, utils.GetLang(language, "about"))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func AboutImgHandler(w http.ResponseWriter, r *http.Request) {
	data := db.GetDocument(utils.AboutImg, "d04b58ce-52d4-4da4-a18b-9820689205ed")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
