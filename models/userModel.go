package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID             uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name           string `json:"name"`
	Password       string `json:"password"`
	NativeLanguage string `json:"native_language"`
}
