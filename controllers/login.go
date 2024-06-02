package controllers

import (
	"net/http"
	"task_web_app/services"

	"github.com/gin-gonic/gin"
)

func ShowSignUpPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.tmpl", gin.H{"Title": "サインアップ"})
}

func PerformSignUp(c *gin.Context) {
	var request struct {
		Email    string `form:"email" binding:"required,email"`
		Password string `form:"password" binding:"required"`
	}

	if err := c.ShouldBind(&request); err != nil {
		c.HTML(http.StatusBadRequest, "signup.tmpl", gin.H{
			"Title": "サインアップ",
			"Error": "全てのフィールドを正しく入力してください。",
		})
		return
	}

	sessionID, err := services.SignUp(request.Email, request.Password)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "signup.tmpl", gin.H{
			"Title": "サインアップ",
			"Error": err.Error(),
		})
	}

	c.SetCookie("session_id", sessionID, 3600, "/", "", true, true)
	c.Redirect(http.StatusFound, "/")
}
