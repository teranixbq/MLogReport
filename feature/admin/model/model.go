package model

import (
	user "mlogreport/feature/user/model"
)

type Admin struct {
	Id       string `gorm:"primaryKey"`
	Role     string `gorm:"type:role_type;"`
	Name     string
	Password string
	Advisor  []user.Users `gorm:"many2many:AdvisorCollege;"`
}

type AdvisorCollege struct {
	AdminsNip string
	UsersNim  string
}
