package services

import (
	"errors"
	"task_web_app/db"
	"task_web_app/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SignUp(email, password string) (string, error) {
	// 登録済みのemailなら登録させない
	existsUser, err := models.FindUserByEmail(email)
	if err != nil {
		return "", err
	}
	if existsUser != nil {
		return "", errors.New("登録済みのemailです")
	}

	// パスワードをハッシュ化してDB漏洩時のユーザー影響を減らす
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	sessionID := uuid.NewString()
	if err := db.Instance().Transaction(func(tx *gorm.DB) error {
		// ユーザーをDBに保存
		newUser, err := models.CreateUser(tx, email, string(hashedPassword))
		if err != nil {
			return err
		}

		// セッションをDBに保存
		return models.CreateSession(tx, newUser.ID, sessionID)
	}); err != nil {
		return "", err
	}

	return sessionID, nil
}

func Login(email, password string) (string, error) {
	existsUser, err := models.FindUserByEmail(email)
	if err != nil {
		return "", err
	}
	if existsUser == nil {
		return "", errors.New("emailかpasswordが間違っています")
	}

	// CompareHashAndPasswordはパスワードが一致しない場合はerrが返される
	if err := bcrypt.CompareHashAndPassword([]byte(existsUser.Password), []byte(password)); err != nil {
		return "", errors.New("emailかpasswordが間違っています")
	}

	sessionID := uuid.NewString()
	if err := db.Instance().Transaction(func(tx *gorm.DB) error {
		// セッションをDBに保存
		return models.CreateSession(tx, existsUser.ID, sessionID)
	}); err != nil {
		return "", err
	}

	return sessionID, nil
}
