package model

import (

	user "mlogreport/feature/user/model"
)

type Admin struct {
	Nip      string   `gorm:"primaryKey"`
	Role     string   `gorm:"type:role_type;"`
	Name     string
	Password string
	Advisor  []user.Users `gorm:"many2many:advisor_college;"`
}
