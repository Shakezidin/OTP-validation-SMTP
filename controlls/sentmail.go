package controlls

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"strconv"
)

func generateOTP() int {
	return rand.Intn(9999 - 1000)
}

func sentMAil(email, name string, otp int) {
	msg := "hello " + name + " your otp is " + strconv.Itoa(otp)
	smtpEmail := os.Getenv("EMAIL")
	smtpPass := os.Getenv("PASSWORD")
	auth := smtp.PlainAuth(
		"",
		smtpEmail,
		smtpPass,
		"smtp.gmail.com",
	)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		smtpEmail,
		[]string{email},
		[]byte(msg),
	)
	if err != nil {
		fmt.Println("error")
	}
}
