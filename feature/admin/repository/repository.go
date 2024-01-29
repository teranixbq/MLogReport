package repository

import (
	"mlogreport/feature/admin/dto/request"

	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

type AdminRepositoryInterface interface {
	CreateAdvisor(data request.CreateAdvisor) error
}

func NewPromptRepository(db *gorm.DB) AdminRepositoryInterface {
	return &adminRepository{
		db: db,
	}
}

func (admin *adminRepository) CreateAdvisor(data request.CreateAdvisor) error {
	input := request.CreateAdvisorToModel(data)

	tx := admin.db.Create(&input)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}