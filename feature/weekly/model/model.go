package model

import "time"

type Weekly struct {
	Id          string
	UsersId     string
	Description string
	CreatedAt   time.Time
}

func (Weekly) TableName() string {
	return "weekly"
}
