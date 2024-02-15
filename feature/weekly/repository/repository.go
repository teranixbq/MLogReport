package repository

import (
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
	UpdateWeekly(id string, data request.RequestWeekly) error
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

func (weekly *weeklyRepository) UpdateWeekly(id string, data request.RequestWeekly) error {

	tx := weekly.db.Model(&model.Weekly{}).Where("id = ? AND users_id = ?", id, data.UsersId).Update("description",data.Description)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
