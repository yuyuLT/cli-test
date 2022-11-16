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

var ideas []Idea

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

func GetIdeas()([]Idea, error){
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
	  panic("Failed to connect database")
	  return nil, nil
	}
	err = db.Limit(5).Order("created_at desc").Find(&ideas).Error
	
	return ideas,err
}