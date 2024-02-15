package service

import (
	"errors"
	"mlogreport/feature/weekly/dto/request"
	"mlogreport/feature/weekly/dto/response"
	"mlogreport/feature/weekly/repository"
	"mlogreport/utils/validation"
	"time"
)

type weeklyService struct {
	weeklyRepository repository.WeeklyRepositoryInterface
}

type WeeklyServiceInterface interface {
	Insert(nim string, data request.RequestWeekly) error
	SelectAll(nim string) ([]response.ResponseWeekly, error)
	UpdateWeekly(id string, data request.RequestWeekly) error
}

func NewWeeklyService(weeklyRepository repository.WeeklyRepositoryInterface) WeeklyServiceInterface {
	return &weeklyService{
		weeklyRepository: weeklyRepository,
	}
}

func (weekly *weeklyService) Insert(nim string, data request.RequestWeekly) error {

	timeAsia, errTime := time.LoadLocation("Asia/Bangkok")
	if errTime != nil {
		return errTime
	}
	day := time.Now().In(timeAsia)

	if day.After(time.Now().In(timeAsia)) {
		return errors.New("error : waktu tidak sesuai")
	}

	errLimit := validation.LimitDescription(data.Description, 5000)
	if errLimit != nil {
		return errLimit
	}

	err := weekly.weeklyRepository.Insert(nim, data)
	if err != nil {
		return err
	}

	return nil
}

func (weekly *weeklyService) SelectAll(nim string) ([]response.ResponseWeekly, error) {
	dataWeekly, err := weekly.weeklyRepository.SelectAll(nim)
	if err != nil {
		return nil, err
	}

	return dataWeekly, nil
}

func (weekly *weeklyService) UpdateWeekly(id string, data request.RequestWeekly) error {

	errLimit := validation.LimitDescription(data.Description, 5000)
	if errLimit != nil {
		return errLimit
	}

	err := weekly.weeklyRepository.UpdateWeekly(id, data)
	if err != nil {
		return err
	}

	return nil
}
