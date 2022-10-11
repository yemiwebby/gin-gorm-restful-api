package model

import (
	"diary_api/database"

	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Content string `gorm:"type:text" json:"content"`
	UserID  uint
}

func (entry *Entry) Save() (*Entry, error) {

	err := database.Database.Create(&entry).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}