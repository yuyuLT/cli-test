package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Idea struct {
	gorm.Model
	Email string
	Date  string
	Text  string
}

var ideas []Idea

func CreateTable() bool {
	db, err := gorm.Open(sqlite.Open("new.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
		return false
	}

	db.AutoMigrate(&Idea{})
	return true
}

func RegisterDataBase(email string, date string, text string) bool {

	db, err := gorm.Open(sqlite.Open("new.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to register database")
		return false
	}

	db.Select("Email", "Date", "Text").Create(&Idea{Email: email, Date: date, Text: text})

	return true
}

func GetIdeas(email string) ([]Idea, error) {
	db, err := gorm.Open(sqlite.Open("new.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
		return nil, nil
	}

	err = db.Where("email = ?", email).Clauses(clause.OrderBy{Expression: clause.Expr{SQL: "RANDOM()", WithoutParentheses: true}}).Limit(5).Find(&ideas).Error

	return ideas, err
}
