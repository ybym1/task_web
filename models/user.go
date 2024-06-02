package models

import (
	"task_web_app/db"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func CreateUser(tx *gorm.DB, email, hashedPassword string) (*User, error) {
	newUser := &User{
		Email:    email,
		Password: hashedPassword,
	}
	if err := tx.Create(newUser).Error; err != nil {
		return nil, err
	}

	return newUser, nil
}

func FindUserByEmail(email string) (*User, error) {
	var existsUser *User
	if err := db.Instance().Where("email = ?", email).First(&existsUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return existsUser, nil
}
