package repository

import (
	admin "mlogreport/feature/admin/model"
	user "mlogreport/feature/user/model"
	"mlogreport/feature/weekly/dto/request"
	"mlogreport/feature/weekly/dto/response"
	"mlogreport/feature/weekly/model"
	"mlogreport/utils/enum"

	"gorm.io/gorm"
)

type weeklyRepository struct {
	db *gorm.DB
}

type WeeklyRepositoryInterface interface {
	Insert(nim string, data request.RequestWeekly) error
	SelectAll(nim string) ([]response.ResponseWeekly, error)
	SelectAllWeeklyAdvisor(nip string) ([]response.ResponseWeekly, error) 
	UpdateWeekly(id string, data request.RequestWeekly) error
	SelectWeekly(id string) (response.ResponseWeekly, error)
	UpdateStatus(id string, status string) error
}

func NewWeeklyRepository(db *gorm.DB) WeeklyRepositoryInterface {
	return &weeklyRepository{
		db: db,
	}
}

func (weekly *weeklyRepository) Insert(nim string, data request.RequestWeekly) error {
	request := request.RequestWeeklyToModel(data)

	request.UsersId = nim
	request.Status = enum.WeeklyStatus[0]
	tx := weekly.db.Create(&request)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (weekly *weeklyRepository) SelectAll(nim string) ([]response.ResponseWeekly, error) {
	dataWeekly := []model.Weekly{}

	tx := weekly.db.Where("users_id = ?", nim).Find(&dataWeekly)
	if tx.Error != nil {
		return nil, tx.Error
	}

	response := response.ListModelToResponseWeekly(dataWeekly)
	return response, nil
}

func (weekly *weeklyRepository) SelectAllWeeklyAdvisor(nip string) ([]response.ResponseWeekly, error) {
	dataAdmin := admin.Admin{}
	dataWeekly := []response.ResponseWeekly{}
	dataUsers := []user.Users{}

	tx := weekly.db.Where("id = ?", nip).First(&dataAdmin)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if txUser := weekly.db.Model(&dataAdmin).Association("Advisor").Find(&dataUsers); txUser != nil {
		return nil, txUser
	}

	for _, user := range dataUsers {
		var userWeekly model.Weekly
		err := weekly.db.Where("users_id = ?", user.Id).First(&userWeekly)

		if err.Error != nil {
			continue
		}

		result := response.ModelToResponseWeeklyAdvisor(userWeekly)
		result.Name = user.Name
		dataWeekly = append(dataWeekly, result)

	}

	return dataWeekly, nil
}

func (weekly *weeklyRepository) UpdateWeekly(id string, data request.RequestWeekly) error {

	tx := weekly.db.Model(&model.Weekly{}).Where("id = ? AND users_id = ?", id, data.UsersId).Update("description", data.Description)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (weekly *weeklyRepository) SelectWeekly(id string) (response.ResponseWeekly, error) {
	dataWeekly := model.Weekly{}

	tx := weekly.db.Where("id = ? ", id).First(&dataWeekly)
	if tx.Error != nil {
		return response.ResponseWeekly{}, nil
	}

	response := response.ModelToResponseWeekly(dataWeekly)
	return response, nil
}

func (weekly *weeklyRepository) UpdateStatus(id string, status string) error {

	tx := weekly.db.Where("id = ? ", id).Update("status", status)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
