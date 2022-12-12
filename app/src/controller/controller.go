package controller

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
	"mvc_test/model"
	"mvc_test/config"
	
)

func GoogleLogin(c *gin.Context) {
	googleConfig := config.SetupConfig();
	url := googleConfig.AuthCodeURL("state")
	c.Redirect(http.StatusSeeOther, url)
}

func GoogleCallback(c *gin.Context) {
}

func ShowTopPage(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{"topPage": "idea_pot"})
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
