package services

import (
	"task_web_app/db"
	"task_web_app/models"

	"gorm.io/gorm"
)

func GetTask(id uint) (*models.Task, error) {
	return models.FindTask(id)
}

func GetTasksByUserID(userID uint) ([]*models.Task, error) {
	return models.FindTasksByUserID(userID)
}

func CreateTask(userID uint, title string, description string) error {
	return db.Instance().Transaction(func(tx *gorm.DB) error {
		return models.CreateTask(tx, userID, title, description)
	})
}
