package models

import (
	"task_web_app/db"

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

func FindSessionByToken(token string) (*Session, error) {
	var existsSession *Session
	if err := db.Instance().Where("token = ?", token).First(&existsSession).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return existsSession, nil
}
