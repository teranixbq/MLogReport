package model

import (
	user "mlogreport/feature/user/model"
)

type Admin struct {
	Nip      string `gorm:"primaryKey"`
	Role     string 
	Name     string
	Password string
	Advisor  []user.Users `gorm:"many2many:advisor_college;"`
}

// func (A *Admin) BeforeCreate(tx *gorm.DB) (err error) {
// 	newUuid := uuid.New()
// 	A.Id = newUuid.String()

// 	return nil
// }
