package models

import "gorm.io/gorm"

type Todo struct {
	ActivityID string
	Title      string
	IsActive   string
	Priority   string
	gorm.Model
}
