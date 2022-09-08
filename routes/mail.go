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
	fmt.Fprintln(w, response)
}
