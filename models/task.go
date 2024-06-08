package models

import (
	"task_web_app/db"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	UserID      uint   `gorm:"not null"`
}

func FindTasksByUserID(userID uint) ([]*Task, error) {
	var tasks []*Task
	if err := db.Instance().Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}