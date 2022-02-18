package main

import (
	"fmt"
	"net/smtp"
	_ "github.com/gin-gonic/gin"
	"os"
)

func init() {
	os.Setenv("FromEmailAddr", "")
	os.Setenv("SMTPpwd", "")
	os.Setenv("ToEmailAddr", "")
}


func main() {
	from := os.Getenv("FromEmailAddr")
	password := os.Getenv("SMTPpwd")

	toEmail := os.Getenv("ToEmailAddr")
	to := []string{toEmail}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	// message
	subject := "Hello"
	body := "first message"
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("go check your email")
}
