package main

import (
	"task_web_app/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controllers.ShowIndexPage)
	r.GET("/signup", controllers.ShowSignUpPage)
	r.Run()
}
