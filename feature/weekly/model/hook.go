package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (w *Weekly) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	w.Id = newUuid.String()

	return nil
}