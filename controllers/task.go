package controllers

import (
	"net/http"
	"task_web_app/models"
	"task_web_app/services"

	"github.com/gin-gonic/gin"
)

func ShowTaskPage(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	tasks, err := services.GetTasksByUserID(user.ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.HTML(http.StatusOK, "tasks.tmpl", gin.H{
		"Title": "タスク一覧",
		"Tasks": tasks,
		"User":  user,
	})
}

func ShowNewTaskPage(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	c.HTML(http.StatusOK, "new_task.tmpl", gin.H{
		"Title": "新しいタスクを作成",
		"User":  user,
	})
}

func CreateTask(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	var request struct {
		Title       string `form:"title" binding:"required"`
		Description string `form:"description" binding:"required"`
	}

	if err := c.ShouldBind(&request); err != nil {
		c.HTML(http.StatusBadRequest, "new_task.tmpl", gin.H{
			"Title": "新しいタスクを作成",
			"User":  user,
			"Error": "入力内容に誤りがあります。",
		})
		return
	}

	if err := services.CreateTask(user.ID, request.Title, request.Description); err != nil {
		c.Error(err)
		return
	}

	c.Redirect(http.StatusFound, "/tasks")
}
