package service

import (
	"mlogreport/feature/weekly/dto/request"
	"mlogreport/feature/weekly/repository"
)

type weeklyService struct {
	weeklyRepository repository.WeeklyRepositoryInterface
}


type WeeklyServiceInterface interface {
	Insert(nim string, data request.RequestWeekly) error
}

func NewWeeklyService(weeklyRepository repository.WeeklyRepositoryInterface) WeeklyServiceInterface {
	return &weeklyService{
		weeklyRepository: weeklyRepository,
	}
}

func (weekly *weeklyService) Insert(nim string, data request.RequestWeekly) error {
	
	err := weekly.weeklyRepository.Insert(nim,data)
	if err != nil {
		return err
	}

	return nil
}