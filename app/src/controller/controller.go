package controller

import (
	"time"
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
	if text == "" {
		text = "初期値"
	}
	today := time.Now()
	const layout = "2006/01/02"
	date := (today.Format(layout))

	//モック
	var list []map[string]string
	list = append(list,map[string]string{ "date" : date, "text" : text})
	list = append(list,map[string]string{ "date" : "2022/09/01", "text" : "アイデアを寝かせる"})
	list = append(list,map[string]string{ "date" : "2021/05/05", "text" : "現代アート"})
	list = append(list,map[string]string{ "date" : "2020/04/02", "text" : "水分を6Lとる"})
	list = append(list,map[string]string{ "date" : "2016/08/20", "text" : "映画の前に大福を食べると良い"})

	c.JSON(200, list)
	
}
