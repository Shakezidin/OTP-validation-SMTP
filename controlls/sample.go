package controlls

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/Shakezidin/model"
	"github.com/gin-gonic/gin"
)

func Sample(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	sentMAil(user.Email, user.Name)
	c.JSON(200, gin.H{"status": "success"})
}

func sentMAil(email, name string) {
	msg := "hello " + name + " ,How are you"
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
