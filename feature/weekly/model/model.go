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

type Periode struct {
	Id        string
	Start     string
	End       string
	CreatedAt time.Time
}

func (Periode) TableName() string {
	return "periode"
}

func (Weekly) TableName() string {
	return "weekly"
}
