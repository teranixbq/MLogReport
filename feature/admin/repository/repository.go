package repository

import (
	"errors"
	"mlogreport/feature/admin/dto/request"
	"mlogreport/feature/admin/dto/response"
	"mlogreport/feature/admin/model"
	user "mlogreport/feature/user/model"
	"mlogreport/utils/enum"
	"mlogreport/utils/meta"

	"github.com/jackc/pgx/v5/pgconn"

	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

type AdminRepositoryInterface interface {
	CreateAdvisor(data request.CreateAdvisor) error
	SelectNip(nip string) (model.Admins, error)
	SelectAllAdvisor(page, limit int) ([]response.ResponseAllAdvisor, meta.Meta, error)
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
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (admin *adminRepository) SelectNip(nip string) (model.Admins, error) {
	dataAdmin := model.Admins{}

	tx := admin.db.Where("nip = ?", nip).Take(&dataAdmin)
	if tx.Error != nil {
		return model.Admins{}, tx.Error
	}

	return dataAdmin, nil
}

func (admin *adminRepository) SelectAllAdvisor(page, limit int) ([]response.ResponseAllAdvisor, meta.Meta, error) {
	dataAdvisor := []model.Admins{}
	var totalData int64
	offset := (page - 1) * limit

	err := admin.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Admins{}).Where("role = ?", enum.RoleType[1]).Count(&totalData).Error; err != nil {
			return err
		}

		if err := tx.
			Where("role = ?", enum.RoleType[1]).
			Offset(offset).
			Limit(limit).
			Order("name ASC").
			Find(&dataAdvisor).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, meta.Meta{}, err
	}

	metaInfo := meta.MetaInfo(page, limit, int(totalData))
	response := response.ListResponseAllAdvisor(dataAdvisor)
	return response, metaInfo, nil
}

func (admin *adminRepository) InsertList(data request.ListCollege) error {

	dataAdmin := model.Admins{}
	err := admin.db.Where("nip = ?", data.Advisor).First(&dataAdmin).Error
	if err != nil {
		return err
	}

	dataUsers := []user.Users{}
	err = admin.db.Where("nim IN (?)", data.Colleges).Find(&dataUsers).Error
	if err != nil {
		return err
	}

	dataManys := make([]model.AdvisorCollege, 0, len(dataUsers))
	for _, user := range dataUsers {
		dataManys = append(dataManys, model.AdvisorCollege{
			UsersId:  user.Id,
			AdminsId: dataAdmin.Id,
		})
	}

	tx := admin.db.Create(&dataManys)

	if errors.As(tx.Error, &pg) {
		return errors.New("error : data already exists")
	}

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (admin *adminRepository) DeleteAdvisor(id string) error {
	dataAdmin := model.Admins{}

	tx := admin.db.Where("id = ? ", id).Delete(&dataAdmin)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
