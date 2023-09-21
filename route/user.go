package route

import (
	"github.com/Shakezidin/controlls"
	"github.com/gin-gonic/gin"
)

func SampleRoute(c *gin.Engine) {
	c.POST("/user", controlls.Sample)
}
