package routes

import (
	"diveEvolution/db"
	"diveEvolution/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	language := params["lang"]
	data := db.GetDocument(utils.Contact, utils.GetLang(language, "contact"))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
func ContactImgHandler(w http.ResponseWriter, r *http.Request) {
	data := db.GetDocument(utils.ContactImg, "3aea6d6c-3fac-4b11-8ee0-5f6a6839da74")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	d, _ := json.Marshal(data)
	w.Write(d)
}
