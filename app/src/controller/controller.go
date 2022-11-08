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
	today := time.Now()
	const layout = "2022/11/25"
	date := (today.Format(layout))

	//モック
	list := []*Item{
		{date, text},
		{"2019/10/22", "おいしい秋刀魚の焼き方"},
		{"2020/04/01", "浜松町絶妙に遠い"},
		{"2022/08/05", "桐麺"},
		{"2022/08/01", "福岡に住みたい"},
	}

	c.JSON(200, list)
}
