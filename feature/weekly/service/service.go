package service

import (
	"errors"
	"time"

	user "mlogreport/feature/user/repository"
	"mlogreport/feature/weekly/dto/request"
	"mlogreport/feature/weekly/dto/response"
	"mlogreport/feature/weekly/repository"
	"mlogreport/utils/constanta"
	"mlogreport/utils/enum"
	"mlogreport/utils/validation"
)

type weeklyService struct {
	weeklyRepository repository.WeeklyRepositoryInterface
	userRepository   user.UserRepositoryInterface
}

type WeeklyServiceInterface interface {
	Insert(nim string, data request.RequestWeekly) error
	SelectAll(nim string) ([]response.ResponseWeekly, error)
	SelectAllWeeklyAdvisor(nip, nim string) (response.ResponseWeeklyDetail, error)
	UpdateWeekly(id string, data request.RequestWeekly) error
	UpdateStatus(idUser, id string, status string) error
	InsertPeriode(data request.RequestPeriode) error
	SelectAllPeriode() ([]response.ResponsePeriode, error)
	UpdatePeriode(id string, data request.RequestPeriode) error
}

func NewWeeklyService(weeklyRepository repository.WeeklyRepositoryInterface, userRepository user.UserRepositoryInterface) WeeklyServiceInterface {
	return &weeklyService{
		weeklyRepository: weeklyRepository,
		userRepository:   userRepository,
	}
}

func (weekly *weeklyService) InsertPeriode(data request.RequestPeriode) error {
	errEmpty := validation.CheckEmpty(data)
	if errEmpty != nil {
		return errEmpty
	}

	errTime := validation.CheckTime(data.Start, data.End, true)
	if errTime != nil {
		return errTime
	}

	dataPeriode, _ := weekly.weeklyRepository.SelectPeriode("")

	if dataPeriode.End != "" {
		errTimeExists := validation.CheckTimeEnd(dataPeriode.End)
		if errTimeExists != nil {
			return errTimeExists
		}
	}

	err := weekly.weeklyRepository.InsertPeriode(data)
	if err != nil {
		return err
	}

	return nil
}

func (weekly *weeklyService) SelectAllPeriode() ([]response.ResponsePeriode, error) {
	dataPeriode, err := weekly.weeklyRepository.SelectAllPeriode()
	if err != nil {
		return nil, err
	}

	return dataPeriode, nil
}

func (weekly *weeklyService) UpdatePeriode(id string, data request.RequestPeriode) error {
	errEmpty := validation.CheckEmpty(data)
	if errEmpty != nil {
		return errEmpty
	}

	errTime := validation.CheckTime(data.Start, data.End, false)
	if errTime != nil {
		return errTime
	}

	dataPeriode, _ := weekly.weeklyRepository.SelectPeriode("")
	_, err := weekly.weeklyRepository.SelectPeriode(id)
	if err != nil {
		return err
	}

	if dataPeriode.Id == "" {
		return errors.New("error : there is no period data yet")
	}

	if dataPeriode.Id != id {
		return errors.New("error : cannot edit old periods")
	}

	err = weekly.weeklyRepository.UpdatePeriode(id, data)
	if err != nil {
		return err
	}

	return nil
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

func (weekly *weeklyService) SelectAllWeeklyAdvisor(nip, nim string) (response.ResponseWeeklyDetail, error) {
	dataWeekly, err := weekly.weeklyRepository.SelectAllWeeklyAdvisor(nip, nim)
	if err != nil {
		return response.ResponseWeeklyDetail{}, err
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

func (weekly *weeklyService) UpdateStatus(idUser, id string, status string) error {
	errEmpty := validation.CheckEmpty(status)
	if errEmpty != nil {
		return errEmpty
	}

	lowerStatus, errValid := validation.CheckEqual(status, enum.WeeklyStatusReq)
	if errValid != nil {
		return errors.New("error : status not match")
	}

	_, err := weekly.userRepository.SelectUserById(idUser)
	if err != nil {
		return err
	}

	dataWeekly, err := weekly.weeklyRepository.SelectWeekly(id)
	if err != nil {
		return err
	}

	if dataWeekly.Status == lowerStatus {
		return errors.New("error : cant edit data again")
	}

	if dataWeekly.Status == constanta.APPROVE {
		return errors.New("error : cannot edit approved data")
	}

	if dataWeekly.Status == constanta.REJECTED && lowerStatus == constanta.APPROVE {
		return errors.New("error : cant edit data again")
	}

	err = weekly.weeklyRepository.UpdateStatus(id, lowerStatus)
	if err != nil {
		return err
	}

	return nil
}
