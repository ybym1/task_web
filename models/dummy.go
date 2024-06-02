package models

import "gorm.io/gorm"

type Dummy struct {
	gorm.Model
	Hoge uint   `gorm:"not null"`
	Fuga string `gorm:"not null"`
}
