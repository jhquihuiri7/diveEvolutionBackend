package utils

import (
	"fmt"
	"net/smtp"
)

type MailRequest struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Subject string `json:"subject"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Message string `json:"message"`
}
func SendMail(data MailRequest) string {
	// sender data
	from := "jhonatan.quihuiri@gmail.com" //ex: "John.Doe@gmail.com"
	password := "dfouvrynuvwkydpi"           // ex: "ieiemcjdkejspqz"
	// receiver address
	toEmail := "jhonatan.quihuiri@gmail.com" // ex: "Jane.Smith@yahoo.com"
	to := []string{toEmail}
	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	// message
	data.Subject = "CONSULTA INFORMACION CLIENTE\r\n"
	message := []byte("Subject:" + data.Subject + "\r\n" +
		"Cliente: " + data.Firstname + " " + data.Lastname + "\r\n" +
		"Mail: " + data.Email + "\r\n" +
		"Celular: " + data.Phone + "\r\n" +
		"Mensaje: " + data.Message + "\r\n")
	// athentication data
	// func PlainAuth(identity, username, password, host string) Auth
	auth := smtp.PlainAuth("", from, password, host)
	// send mail
	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println("err:", err)
		return "no"
	}
	AutomaticResponse(data.Email)
	return "ok"
}
func AutomaticResponse(toEmail string)  {
	// sender data
	from := "jhonatan.quihuiri@gmail.com" //ex: "John.Doe@gmail.com"
	password := "dfouvrynuvwkydpi"           // ex: "ieiemcjdkejspqz"
	// receiver address
	to := []string{toEmail}
	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	// message
	subject := "DIVE EVOLUTION CONTACTO\n"
	body := "Hemos recibido su mensaje, nos prondremos en contacto pronto"
	message := []byte(subject + body)
	// athentication data
	// func PlainAuth(identity, username, password, host string) Auth
	auth := smtp.PlainAuth("", from, password, host)
	// send mail
	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}
