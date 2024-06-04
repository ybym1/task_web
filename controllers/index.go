package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	user, _ := c.Get("user")
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"Title": "ホーム", "User": user})
}

func Show500Page(c *gin.Context) {
	c.HTML(http.StatusInternalServerError, "500.tmpl", gin.H{
		"Title": "内部サーバーエラー",
	})
}
