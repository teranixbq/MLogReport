package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	u.Id = newUuid.String()
	return nil
}
