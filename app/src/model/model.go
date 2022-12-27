package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Idea struct {
	gorm.Model
	UserId int
	Date   string
	Text   string
}

var ideas []Idea

func CreateTable() bool {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
		return false
	}

	db.AutoMigrate(&Idea{})
	return true
}

func RegisterDataBase(date string, text string) bool {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
		return false
	}

	db.Select("UserId", "Date", "Text").Create(&Idea{UserId: 1, Date: date, Text: text})

	return true
}

func GetIdeas() ([]Idea, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
		return nil, nil
	}
	err = db.Clauses(clause.OrderBy{Expression: clause.Expr{SQL: "RANDOM()", WithoutParentheses: true}}).Limit(5).Find(&ideas).Error

	return ideas, err
}
