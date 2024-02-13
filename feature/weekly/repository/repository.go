package repository

import (
	"mlogreport/feature/weekly/model"
	"mlogreport/feature/weekly/dto/request"
	"mlogreport/feature/weekly/dto/response"

	"gorm.io/gorm"
)

type weeklyRepository struct {
	db *gorm.DB
}

type WeeklyRepositoryInterface interface {
	Insert(nim string, data request.RequestWeekly) error
	SelectAll(nim string) ([]response.ResponseWeekly, error)
}

func NewWeeklyRepository(db *gorm.DB) WeeklyRepositoryInterface {
	return &weeklyRepository{
		db: db,
	}
}

func (weekly *weeklyRepository) Insert(nim string, data request.RequestWeekly) error {
	request := request.RequestWeeklyToModel(data)

	request.UsersId = nim
	tx := weekly.db.Create(&request)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (weekly *weeklyRepository) SelectAll(nim string) ([]response.ResponseWeekly, error) {
	dataWeekly := []model.Weekly{}

	tx := weekly.db.Where("users_id = ?",nim).Find(&dataWeekly)
	if tx.Error != nil {
		return nil,tx.Error
	}

	response := response.ListModelToResponseWeekly(dataWeekly)
	return response,nil
}