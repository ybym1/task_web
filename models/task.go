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

func FindTask(id uint) (*Task, error) {
	var task *Task
	if err := db.Instance().First(&task, id).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func FindTasksByUserID(userID uint) ([]*Task, error) {
	var tasks []*Task
	if err := db.Instance().Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func CreateTask(tx *gorm.DB, userID uint, title, description string) error {
	return tx.Create(&Task{
		UserID:      userID,
		Title:       title,
		Description: description,
	}).Error
}
