package repository

import (
	"errors"
	"mlogreport/feature/admin/dto/request"
	"mlogreport/feature/admin/model"

	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

type AdminRepositoryInterface interface {
	CreateAdvisor(data request.CreateAdvisor) error
	FindNip(nip string) (model.Admin, error)
}

func NewPromptRepository(db *gorm.DB) AdminRepositoryInterface {
	return &adminRepository{
		db: db,
	}
}

func (admin *adminRepository) CreateAdvisor(data request.CreateAdvisor) error {
	input := request.CreateAdvisorToModel(data)

	tx := admin.db.Create(&input)

	if tx.RowsAffected == 0 {
		return errors.New("error : nip already exists")
	}

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (admin *adminRepository) FindNip(nip string) (model.Admin, error) {
	dataAdmin := model.Admin{}

	tx := admin.db.Where("nip = ?", nip).First(&dataAdmin)
	if tx.Error != nil {
		return model.Admin{}, tx.Error
	}

	return dataAdmin, nil
}
