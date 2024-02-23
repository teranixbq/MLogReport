package model

import (
	user "mlogreport/feature/user/model"

	"gorm.io/gorm"
)

type Admins struct {
	Id        string `gorm:"primaryKey"`
	Nip       string
	Role      string `gorm:"type:role_type;"`
	Name      string
	Password  string
	Advisor   []user.Users   `gorm:"many2many:AdvisorCollege;"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type AdvisorCollege struct {
	AdminsId string
	UsersId  string
}
