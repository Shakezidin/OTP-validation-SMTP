package main

import (
	"github.com/Shakezidin/initializer"
	"github.com/Shakezidin/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	initializer.Load()
	route.SampleRoute(r)
	r.Run("localhost:8080")
}
