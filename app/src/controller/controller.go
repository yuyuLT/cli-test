package controller

import (
	"github.com/gin-gonic/gin"
	oauthapi "google.golang.org/api/oauth2/v2"
	"time"
	"net/http"
	"context"
	"errors"
	"fmt"
	"mvc_test/model"
	"mvc_test/config"
)

func ShowTop(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"name": "name",
	})
}

func GoogleLogin(c *gin.Context) {
	googleConfig := config.SetupConfig()
	url := googleConfig.AuthCodeURL("state")
	c.Redirect(http.StatusSeeOther, url)
}

func GoogleCallback(c *gin.Context) {
	config := config.SetupConfig()
    context := context.Background()
	code := c.Query("code")

    tok, err := config.Exchange(context, code)
    if err != nil {
        panic(err)
    }

    if tok.Valid() == false {
        panic(errors.New("vaild token"))
    }

    client := config.Client(context, tok)
    svr, err := oauthapi.New(client)
	ui, err := svr.Userinfo.Get().Do()

	if err != nil {
		fmt.Println("OAuth Error")
		panic(err)
	} 

	userName := ui.Name
	userEmail := ui.Email
	fmt.Println(userName)
	fmt.Println(userEmail)

	c.Redirect(http.StatusFound, "/top")
	
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
