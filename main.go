package main

import (
	"errors"
	"net/http"
	"task_web_app/controllers"
	"task_web_app/db"
	"task_web_app/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Instance().AutoMigrate(&models.Dummy{}, &models.User{}, &models.Session{})

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Use(entry)
	r.GET("/", controllers.ShowIndexPage)
	r.GET("/signup", controllers.ShowSignUpPage)
	r.POST("/signup", controllers.PerformSignUp)
	r.GET("/login", controllers.ShowLoginPage)
	r.POST("/login", controllers.PerformLogin)
	r.Run()
}

func entry(c *gin.Context) {
	if err := setupSession(c); err != nil {
		c.Abort()
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Next()
}

func setupSession(c *gin.Context) error {
	sessionID, err := c.Cookie("session_id")
	if err != nil && err != http.ErrNoCookie {
		return errors.Join(errors.New("セッションの取得に失敗しました"), err)
	}

	if sessionID == "" {
		return nil
	}

	session, err := models.FindSessionByToken(sessionID)
	if err != nil {
		return errors.Join(errors.New("セッションの取得に失敗しました"), err)
	}

	if session == nil {
		return nil
	}

	user, err := models.FindUserByID(session.UserID)
	if err != nil {
		return errors.Join(errors.New("ユーザーの取得に失敗しました"), err)
	}
	c.Set("user", user)
	return nil
}
