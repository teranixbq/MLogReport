package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *Report) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	r.Id = newUuid.String()

	return nil
}