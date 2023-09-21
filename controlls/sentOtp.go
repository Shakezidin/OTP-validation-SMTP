package controlls

import (
	"github.com/Shakezidin/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SentOtp(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	Otp := generateOTP()

	sentMAil(user.Email, user.Name, Otp)

	session := sessions.Default(c)
	session.Set(user.Email, Otp)
	session.Save()
	c.JSON(200, gin.H{"status": "otp sent success"})
}
