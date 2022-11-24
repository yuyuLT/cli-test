package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("view/*html")
	r.Static("assets", "./assets")
	r.GET("/", ShowTopPage)
	r.POST("/api", Sendtext)
	return r
}