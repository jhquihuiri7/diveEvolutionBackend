package routes

import (
	"diveEvolution/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func MailSender(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Headers:", "*")
	w.Header().Add("Access-Control-Allow-Origin", "Accept-Encoding")
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//, Authorization, Content-Type, Content-Length, Accept-Encoding
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	var data = utils.MailRequest{}
	responseR, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal(responseR, &data)

	//fmt.Println(data)
	err = data.ParseTemplate()
	if err != nil {
		log.Fatalln(err)
	}
	response := data.SendMail()
	JSONresponse, err := json.Marshal(response)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintln(w, string(JSONresponse))
}
