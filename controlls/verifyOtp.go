package controlls

import (
	"github.com/Shakezidin/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func OtpVerification(c *gin.Context) {
	var verify model.Verify
	if err := c.BindJSON(&verify); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	session := sessions.Default(c)
	if session.Get(verify.Email) != verify.OTP {
		c.JSON(400, gin.H{"Error": "otp error"})
		return
	}

	c.JSON(200, gin.H{"status": "otp correct"})
}
