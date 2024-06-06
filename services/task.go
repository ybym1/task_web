package services

import (
	"task_web_app/models"
)

func GetTasksByUserID(userID uint) ([]*models.Task, error) {
	return models.FindTasksByUserID(userID)
}
