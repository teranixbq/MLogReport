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
	SelectAllWeeklyAdvisor(nip, nim string) (response.ResponseWeeklyDetail, error)
	UpdateWeekly(id string, data request.RequestWeekly) error
	SelectWeekly(id string) (response.ResponseWeekly, error)
	UpdateStatus(id string, status string) error
	InsertPeriode(data request.RequestPeriode) error
	SelectPeriode(id string) (response.ResponsePeriode, error)
	SelectAllPeriode() ([]response.ResponsePeriode, error)
	UpdatePeriode(id string, data request.RequestPeriode) error
}

func NewWeeklyRepository(db *gorm.DB) WeeklyRepositoryInterface {
	return &weeklyRepository{
		db: db,
	}
}

func (weekly *weeklyRepository) InsertPeriode(data request.RequestPeriode) error {
	request := request.RequestPeriodeToModel(data)

	tx := weekly.db.Create(&request)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (weekly *weeklyRepository) SelectPeriode(id string) (response.ResponsePeriode, error) {
	dataPeriode := model.Periode{}
	if id == "" {
		tx := weekly.db.Order("created_at DESC").First(&dataPeriode)
		if tx.Error != nil {
			return response.ResponsePeriode{}, tx.Error
		}
	} else {
		tx := weekly.db.Where("id = ? ", id).First(&dataPeriode)
		if tx.Error != nil {
			return response.ResponsePeriode{}, tx.Error
		}
	}

	response := response.ModelToResponsePeriode(dataPeriode)
	return response, nil
}

func (weekly *weeklyRepository) SelectAllPeriode() ([]response.ResponsePeriode, error) {
	dataPeriode := []model.Periode{}

	tx := weekly.db.Order("created_at DESC").Find(&dataPeriode)
	if tx.Error != nil {
		return nil, tx.Error
	}

	response := response.ListModelToResponsePeriode(dataPeriode)
	return response, nil
}

func (weekly *weeklyRepository) UpdatePeriode(id string, data request.RequestPeriode) error {
	request := request.RequestPeriodeToModel(data)

	tx := weekly.db.Where("id = ? ", id).Updates(&request)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (weekly *weeklyRepository) Insert(id string, data request.RequestWeekly) error {
	request := request.RequestWeeklyToModel(data)

	request.UsersId = id
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

func (weekly *weeklyRepository) SelectAllWeeklyAdvisor(IdAdmin, IdUser string) (response.ResponseWeeklyDetail, error) {
	dataAdmin := admin.Admins{}
	dataUser := user.Users{}
	var userWeekly []model.Weekly

	tx := weekly.db.Where("id = ?", IdAdmin).First(&dataAdmin)
	if tx.Error != nil {
		return response.ResponseWeeklyDetail{}, tx.Error
	}

	tx = weekly.db.Where("id = ?", IdUser).First(&dataUser)
	if tx.Error != nil {
		return response.ResponseWeeklyDetail{}, tx.Error
	}

	err := weekly.db.Model(&dataAdmin).Association("Advisor").Find(&dataUser)
	if err != nil {
		return response.ResponseWeeklyDetail{}, err
	}

	tx = weekly.db.Where("users_id = ?", dataUser.Id).Find(&userWeekly)
	if tx.Error != nil {
		return response.ResponseWeeklyDetail{}, tx.Error
	}

	response := response.ModelToResponseWeeklyDetail(dataUser.Nim, dataUser.Name, userWeekly)
	return response, nil
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

	tx := weekly.db.Where("id = ?", id).First(&dataWeekly)
	if tx.Error != nil {
		return response.ResponseWeekly{}, nil
	}

	response := response.ModelToResponseWeekly(dataWeekly)
	return response, nil
}

func (weekly *weeklyRepository) UpdateStatus(id string, status string) error {
	tx := weekly.db.Model(&model.Weekly{}).Where("id = ? ", id).Update("status", status)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
