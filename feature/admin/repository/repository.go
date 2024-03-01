package repository

import (
	"errors"
	"fmt"

	"mlogreport/feature/admin/dto/request"
	"mlogreport/feature/admin/dto/response"
	"mlogreport/feature/admin/model"
	user "mlogreport/feature/user/model"
	"mlogreport/utils/constanta"
	"mlogreport/utils/enum"
	"mlogreport/utils/meta"

	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

type AdminRepositoryInterface interface {
	CreateAdvisor(data request.CreateAdvisor) error
	SelectNip(nip string) (model.Admins, error)
	SelectAllAdvisor(page, limit int) ([]response.ResponseAllAdvisor, meta.Meta, error)
	SelectAdvisor(id string) (response.ResponseAdvisor, error)
	InsertList(data request.ListCollege) error
	DeleteAdvisor(id string) error
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

	tx := admin.db.Model(&model.Admins{}).Where("role = ?", enum.RoleType[1]).Count(&totalData)
	if tx.Error != nil {
		return nil, meta.Meta{}, tx.Error
	}

	tx = admin.db.
		Where("role = ?", enum.RoleType[1]).
		Offset(offset).
		Limit(limit).
		Order("name ASC").
		Find(&dataAdvisor)

	if tx.Error != nil {
		return nil, meta.Meta{}, tx.Error
	}

	metaInfo := meta.MetaInfo(page, limit, int(totalData))
	response := response.ListResponseAllAdvisor(dataAdvisor)
	return response, metaInfo, nil
}

func (admin *adminRepository) SelectAdvisor(id string) (response.ResponseAdvisor, error) {
	dataAdmin := model.Admins{}

	if err := admin.db.Preload("Advisor").First(&dataAdmin, "id = ?", id).Error; err != nil {
		return response.ResponseAdvisor{}, err
	}

	response := response.ModelToResponseAdvisor(dataAdmin)
	return response, nil
}

func (admin *adminRepository) InsertList(data request.ListCollege) error {
	dataAdmin := model.Admins{}

	err := admin.db.Where("nip = ?", data.Advisor).First(&dataAdmin)
	if err.Error != nil {
		return err.Error
	}

	dataUsers := []user.Users{}
	for _, college := range data.Colleges {
		var user user.Users
		var advisorCollege model.AdvisorCollege

		err := admin.db.Where("nim = ?", college).First(&user)
		if err.Error != nil {
			return err.Error
		}

		err = admin.db.Where("users_id = ? ", user.Id).First(&advisorCollege)
		if err.RowsAffected != 0 {
			err := fmt.Sprintf(constanta.EXISTS,user.Nim)
			return errors.New(err)
		}

		dataUsers = append(dataUsers, user)
	}

	dataManys := make([]model.AdvisorCollege, 0, len(dataUsers))
	for _, user := range dataUsers {
		dataManys = append(dataManys, model.AdvisorCollege{
			UsersId:  user.Id,
			AdminsId: dataAdmin.Id,
		})
	}

	err = admin.db.Create(&dataManys)
	if err.Error != nil {
		return err.Error
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
