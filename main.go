package main

import (
	"task_web_app/controllers"
	"task_web_app/db"
	"task_web_app/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Instance().AutoMigrate(&models.Dummy{}, &models.User{}, &models.Session{})

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controllers.ShowIndexPage)
	r.GET("/signup", controllers.ShowSignUpPage)
	r.POST("/signup", controllers.PerformSignUp)
	r.GET("/login", controllers.ShowLoginPage)
	r.POST("/login", controllers.PerformLogin)
	r.Run()
}
