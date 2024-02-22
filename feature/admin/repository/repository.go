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
	FindNip(id string) (model.Admin, error)
	InsertList(data request.ListCollege) error
	DeleteAdvisor(id string) error
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

	tx := admin.db.Where("nip = ?", nip).First(&dataAdmin)
	if tx.Error != nil {
		return model.Admin{}, tx.Error
	}

	return dataAdmin, nil
}

func (admin *adminRepository) InsertList(data request.ListCollege) error {
	adminData := model.Admin{}

	tx := admin.db.Preload("Advisor").Where("nip = ?", data.Advisor).Take(&adminData)
	if tx.Error != nil {
		return tx.Error
	}

	for _, usersNim := range data.Colleges {
		dataUsers:= user.Users{}

		tx := admin.db.Where("nim = ?",usersNim).Take(&dataUsers)
		if tx.Error != nil {
			return tx.Error
		}

		tx = admin.db.Exec("INSERT INTO advisor_colleges (admin_id, users_id) VALUES (?, ?)", adminData.Id, dataUsers.Id)

		if errors.As(tx.Error, &pg) {
			return errors.New("ERROR : data already exists")
		}

		if tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}

func (admin *adminRepository) DeleteAdvisor(id string) error {
	dataAdmin := model.Admin{}
	
	tx := admin.db.Where("id = ? ", id).Delete(&dataAdmin)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
