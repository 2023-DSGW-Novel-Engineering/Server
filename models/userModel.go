package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID             uint `gorm:"primaryKey;autoIncrement"`
	Name           string
	NativeLanguage string
	ImagePath      string
}
