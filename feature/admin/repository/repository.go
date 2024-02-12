package repository

import (
	"errors"
	"mlogreport/feature/admin/dto/request"
	"mlogreport/feature/admin/model"
	user "mlogreport/feature/user/model"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

type AdminRepositoryInterface interface {
	CreateAdvisor(data request.CreateAdvisor) error
	FindNip(nip string) (model.Admin, error)
	InsertList(data request.ListCollege) error
}

func NewPromptRepository(db *gorm.DB) AdminRepositoryInterface {
	return &adminRepository{
		db: db,
	}
}

var (
	pg *pgconn.PgError
)

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

	tx := admin.db.Where("id = ?", nip).First(&dataAdmin)
	if tx.Error != nil {
		return model.Admin{}, tx.Error
	}

	return dataAdmin, nil
}

func (admin *adminRepository) InsertList(data request.ListCollege) error {
	adminData := model.Admin{}

	tx := admin.db.Preload("Advisor").Where("id = ?", data.Advisor).First(&adminData)
	if tx.Error != nil {
		return tx.Error
	}

	for _, usersNim := range data.Colleges {
		dataAdd := user.Users{}

		tx := admin.db.First(&dataAdd, "id = ?", usersNim)
		// tx := admin.db.Exec("INSERT INTO advisor_college (admin_nip, users_nim) VALUES (?, ?)", data.Advisor, usersNim)
		if tx.Error != nil {
			return tx.Error
		}

		adminData.Advisor = append(adminData.Advisor, dataAdd)
	}

	tx = admin.db.Create(&adminData)
	if errors.As(tx.Error, &pg) {
			return errors.New("ERROR : data already exists")
		}
		
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
