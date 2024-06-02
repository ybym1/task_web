package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3306)/task_web?charset=utf8mb4&parseTime=True"

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MySQL database!")
}

func Instance() *gorm.DB {
	return db
}
