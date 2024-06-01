package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"Title": "ホーム"})
}
