package controller

import (
	"time"
	"mvc_test/model"
	"github.com/gin-gonic/gin"
)

func ShowMessage(c *gin.Context) {
	message := model.GetMessage()
	c.HTML(200, "index.html", gin.H{"message": message})
}

func Sendtext(c *gin.Context) {

	today := time.Now()
	const layout = "2006/01/02"
	date := (today.Format(layout))

	text := c.PostForm("text")
	if text == "" {
		text = "初期値"
	}else{
		model.RegisterDataBase(date,text);
	}
	
	ideas, _ := model.GetIdeas()
	
	var list []map[string]string
	for _, v := range ideas{
		list = append(list,map[string]string{ "date" : v.Date, "text" : v.Text})
	}

	c.JSON(200, list)
	
}
