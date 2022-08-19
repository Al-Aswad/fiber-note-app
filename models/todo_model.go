package models

import "gorm.io/gorm"

type Todo struct {
	ActivityGroupID int
	Title           string
	IsActive        bool
	Priority        string
	gorm.Model
}
