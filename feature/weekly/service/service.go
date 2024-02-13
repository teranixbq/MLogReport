package service

import (
	"mlogreport/feature/weekly/dto/request"
	"mlogreport/feature/weekly/dto/response"
	"mlogreport/feature/weekly/repository"
)

type weeklyService struct {
	weeklyRepository repository.WeeklyRepositoryInterface
}

type WeeklyServiceInterface interface {
	Insert(nim string, data request.RequestWeekly) error
	SelectAll(nim string) ([]response.ResponseWeekly, error)
}

func NewWeeklyService(weeklyRepository repository.WeeklyRepositoryInterface) WeeklyServiceInterface {
	return &weeklyService{
		weeklyRepository: weeklyRepository,
	}
}

func (weekly *weeklyService) Insert(nim string, data request.RequestWeekly) error {

	err := weekly.weeklyRepository.Insert(nim, data)
	if err != nil {
		return err
	}

	return nil
}

func (weekly *weeklyService) SelectAll(nim string) ([]response.ResponseWeekly, error) {
	dataWeekly,err := weekly.weeklyRepository.SelectAll(nim)
	if err != nil {
		return nil,err
	}

	return dataWeekly,nil
}