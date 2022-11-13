package model

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Idea struct {
	gorm.Model
	UserId int
	Date string
	Text string
}

func GetMessage()(string){
	message := "変数で入力したメッセージです"
	return message
}

func RegisterDataBase(date string,text string)bool{

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
	  panic("Failed to connect database")
	  return false
	}

	db.AutoMigrate(&Idea{})
	db.Select("UserId", "Date", "Text").Create(&Idea{UserId: 1 , Date: date, Text: text})

	return true
}