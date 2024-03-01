package model

import (
	"time"

	report "mlogreport/feature/report/model"
	weekly "mlogreport/feature/weekly/model"

	"gorm.io/gorm"
)

type Users struct {
	Id        string `gorm:"primaryKey"`
	Nim       string `gorm:"unique;not null"`
	Name      string
	Password  string
	Class     string
	Mitra     string
	Program   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt  `gorm:"index"`
	Weekly    []weekly.Weekly `gorm:"foreignKey:UsersId"`
	Report    report.Report   `gorm:"foreignKey:UsersId"`
}
