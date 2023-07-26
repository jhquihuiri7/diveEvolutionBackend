package utils

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
)

type MailRequest struct {
	Name     string `json:"name"`
	Subject  string `json:"subject"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Message  string `json:"message"`
	Response string `json:"response"`
}

func (r *MailRequest) SendMail() MailStruct {
	//response
	response := MailStruct{}
	// sender data
	from := "jhonatan.quihuiri@gmail.com" //ex: "John.Doe@gmail.com"
	password := "dfouvrynuvwkydpi"        // ex: "ieiemcjdkejspqz"
	// receiver address
	toEmail := "jhonatan.quihuiri@gmail.com" // ex: "Jane.Smith@yahoo.com"
	to := []string{toEmail}
	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	subject := "Subject: NUEVA CONSULTA CLIENTE\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := r.Message
	message := []byte(fmt.Sprintf("From: Dive Evolution <%s>\n", from) + subject + mime + body)
	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println("err:", err)
		response.Error = "Error al enviar mensaje"
	} else {
		response.Success = "Correo enviado exitosamente"
		r.AutomaticResponse()
	}
	return response
}
func (r *MailRequest) AutomaticResponse() string {
	// sender data
	from := "jhonatan.quihuiri@gmail.com" //ex: "John.Doe@gmail.com"
	password := "dfouvrynuvwkydpi"        // ex: "ieiemcjdkejspqz"
	// receiver address
	to := []string{r.Email}
	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	// message
	subject := "Subject: ¡Hemos recibido tu mensaje!\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := r.Response
	message := []byte(fmt.Sprintf("From: Dive Evolution <%s>\n", from) + subject + mime + body)
	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(address, auth, from, to, message)

	if err != nil {
		return err.Error()
	} else {
		return ""
	}
}
func (r *MailRequest) ParseTemplate() error {
	t := template.Must(template.ParseGlob("templates/*.gohtml"))
	bufRequest := new(bytes.Buffer)
	bufResponse := new(bytes.Buffer)
	if err := t.ExecuteTemplate(bufRequest, "requestMail.gohtml", r); err != nil {
		return err
	}
	if err := t.ExecuteTemplate(bufResponse, "responseMail.gohtml", nil); err != nil {
		return err
	}
	r.Message = bufRequest.String()
	r.Response = bufResponse.String()

	return nil
}
