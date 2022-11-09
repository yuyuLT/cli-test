package controller

import (
	"time"
	// "encoding/json"
	"mvc_test/model"
	"github.com/gin-gonic/gin"
)

type Item struct {
	date string
	title  string
}

func ShowMessage(c *gin.Context) {
	message := model.GetMessage()
	c.HTML(200, "index.html", gin.H{"message": message})
}

func Sendtext(c *gin.Context) {

	/*
       DB操作
    */
	text := c.PostForm("text")
	today := time.Now()
	const layout = "2006/01/02"
	date := (today.Format(layout))

	//モック
	var list []map[string]string
	list = append(list,map[string]string{"date":date, "text":text})
	list = append(list,map[string]string{"date":"あああ", "text":"いいい"})

	// json_list, err :=  json.Marshal(list)

	// if err == nil {
		c.JSON(200, list)
	// }
	
}
