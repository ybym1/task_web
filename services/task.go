package services

import (
	"errors"
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

func UpdateTask(id uint, userID uint, title string, description string) error {
	task, err := models.FindTask(id)
	if err != nil {
		return err
	}

	if task.UserID != userID {
		return errors.New("このタスクを編集する権限がありません")
	}

	task.Title = title
	task.Description = description
	return db.Instance().Transaction(func(tx *gorm.DB) error {
		return models.UpdateTask(tx, task)
	})
}

func DeleteTask(id uint, userID uint) error {
	task, err := models.FindTask(id)
	if err != nil {
		return err
	}

	if task.UserID != userID {
		return errors.New("このタスクを編集する権限がありません")
	}

	return db.Instance().Transaction(func(tx *gorm.DB) error {
		return models.DeleteTask(tx, id)
	})
}
