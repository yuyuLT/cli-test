package controller

import (
	"context"
	"errors"
	"fmt"
	"mvc_test/config"
	"mvc_test/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	oauthapi "google.golang.org/api/oauth2/v2"
)

func ShowTop(c *gin.Context) {
	userName, _ := model.GetKey(c, "userName")
	userEmail, _ := model.GetKey(c, "userEmail")
	userPicture, _ := model.GetKey(c, "userPicture")

	c.HTML(200, "index.html", gin.H{
		"userName":    userName,
		"userEmail":   userEmail,
		"userPicture": userPicture,
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
	userPicture := ui.Picture

	model.SetKey(c, "userName", userName)
	model.SetKey(c, "userEmail", userEmail)
	model.SetKey(c, "userPicture", userPicture)

	c.Redirect(http.StatusFound, "/top")

}

func Sendtext(c *gin.Context) {
	email, _ := model.GetKey(c, "userEmail")
	model.CreateTable()

	today := time.Now()
	const layout = "2006/01/02"
	date := (today.Format(layout))

	text := c.PostForm("text")
	if text == "" {
		text = "初期値"
	} else {
		model.RegisterDataBase(email, date, text)
	}

	ideas, _ := model.GetIdeas(email)

	var list []map[string]string
	for _, v := range ideas {
		list = append(list, map[string]string{"date": v.Date, "text": v.Text})
	}

	c.JSON(200, list)

}
