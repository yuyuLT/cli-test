package controller

import (
	"mvc_test/model"
	"github.com/gin-gonic/gin"
)

func ShowMessage(c *gin.Context) {
	message := model.GetMessage()
	c.HTML(200, "index.html", gin.H{"message": message})
}
