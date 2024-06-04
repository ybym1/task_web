package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	user, _ := c.Get("user")
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"Title": "ホーム", "User": user})
}
