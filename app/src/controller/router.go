package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("view/*html")
	r.Static("assets", "./assets")
	r.GET("/", GoogleLogin)
	r.GET("/google/callback", GoogleCallback)
	r.GET("/top", ShowTop)
	r.POST("/api", Sendtext)
	return r
}