package main

import (
	"github.com/Shakezidin/initializer"
	"github.com/Shakezidin/route"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	initializer.Load()
	store := cookie.NewStore([]byte("iamsuperkey"))
	r.Use(sessions.Sessions("mysession", store))
	
	route.SampleRoute(r)
	r.Run("localhost:8080")
}
