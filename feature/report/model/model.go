package model

import "time"

type Report struct {
	Id            string
	UsersId       string
	FinalReport   string
	Transcript    string
	Certification string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (Report) TableName() string {
	return "report"
}
