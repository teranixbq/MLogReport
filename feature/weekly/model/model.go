package model

import "time"

type Weekly struct {
	Id          string
	UsersId     string
	Description string
	Status      string `gorm:"type:weekly_status;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Weekly) TableName() string {
	return "weekly"
}
