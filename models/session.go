package models

import (
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	UserID uint   `gorm:"not null"`
	Token  string `gorm:"unique;not null"`
}

func CreateSession(tx *gorm.DB, userID uint, token string) error {
	return tx.Create(&Session{
		UserID: userID,
		Token:  token,
	}).Error
}
