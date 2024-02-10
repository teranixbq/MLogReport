package model

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Nim         string `gorm:"primaryKey"`
	Name        string
	Password    string
	Class       string
	Program     string
	Total_Score float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
