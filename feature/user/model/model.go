package model

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Nim       string `gorm:"primaryKey"`
	Name      string
	Password  string
	Class     string
	Mitra     string
	Program   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
