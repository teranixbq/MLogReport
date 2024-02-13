package model

import (
	"time"

	weekly "mlogreport/feature/weekly/model"

	"gorm.io/gorm"
)

type Users struct {
	Id        string `gorm:"primaryKey"`
	Name      string
	Password  string
	Class     string
	Mitra     string
	Program   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt  `gorm:"index"`
	Weekly    []weekly.Weekly `gorm:"foreignKey:UsersId"`
}
