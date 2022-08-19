package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID        uint           `json:"id"`
	Email     string         `json:"email"`
	Title     string         `json:"title"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
