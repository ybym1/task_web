package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowSignUpPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.tmpl", gin.H{"Title": "サインアップ"})
}
