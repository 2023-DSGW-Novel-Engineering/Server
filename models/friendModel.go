package models

type Friend struct {
	Me     int `gorm:"primaryKey"`
	Target int `gorm:"primaryKey"`
}
