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
	r.GET("/500", controllers.Show500Page)
	r.GET("/", controllers.ShowIndexPage)
	r.GET("/signup", controllers.ShowSignUpPage)
	r.POST("/signup", controllers.PerformSignUp)
	r.GET("/login", controllers.ShowLoginPage)
	r.POST("/login", controllers.PerformLogin)
	r.Run()
}

func entry(c *gin.Context) {
	// 無限エラーループ防止のため500ページ表示時は極力何もしない
	if c.Request.URL.Path == "/500" {
		c.Next()
		return
	}

	if err := setupSession(c); err != nil {
		// セッション処理中にエラーがあれば500ページにリダイレクトする
		c.Error(err)
		c.Abort()
		c.Redirect(http.StatusFound, "/500")
		return
	}

	c.Next()

	// controller以下の処理中にエラーがあれば500ページにリダイレクトする
	if len(c.Errors) > 0 {
		c.Redirect(http.StatusFound, "/500")
	}
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
