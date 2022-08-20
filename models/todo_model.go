package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID              uint           `json:"id"`
	ActivityGroupID int            `json:"activity_group_id"`
	Title           string         `json:"title"`
	IsActive        bool           `json:"is_active"`
	Priority        string         `json:"priority"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"update_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
