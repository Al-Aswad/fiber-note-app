package models

import "gorm.io/gorm"

type Activity struct {
	Email string
	Title string
	Todo  []Todo `gorm:"foreignKey:ID;references:ID"`
	gorm.Model
}
