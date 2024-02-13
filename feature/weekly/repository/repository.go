package repository

import (
	"mlogreport/feature/weekly/dto/request"

	"gorm.io/gorm"
)

type weeklyRepository struct {
	db *gorm.DB
}

type WeeklyRepositoryInterface interface {
	Insert(nim string,data request.RequestWeekly) error
}

func NewWeeklyRepository(db *gorm.DB) WeeklyRepositoryInterface {
	return &weeklyRepository{
		db: db,
	}
}

func (weekly *weeklyRepository) Insert(nim string,data request.RequestWeekly) error {
	request := request.RequestWeeklyToModel(data)

	request.UsersId = nim
	tx := weekly.db.Create(&request)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}