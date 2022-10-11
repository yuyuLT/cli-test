package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("view/*html")
	r.Static("images", "./images")
	r.GET("/", ShowMessage)
	return r
}