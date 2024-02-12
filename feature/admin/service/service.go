package service

import (
	"errors"
	"mlogreport/feature/admin/dto/request"
	"mlogreport/feature/admin/dto/response"
	"mlogreport/feature/admin/repository"
	"mlogreport/utils/auth"
	"mlogreport/utils/constanta"
	"mlogreport/utils/helper"
	"mlogreport/utils/validation"
)

type adminService struct {
	adminRepository repository.AdminRepositoryInterface
}

type AdminServiceInterface interface {
	CreateAdvisor(data request.CreateAdvisor) error
	Login(data request.AdminLogin) (response.ResponseLogin, error)
	InsertList(data request.ListCollege) error
}

func NewAdminService(adminRepository repository.AdminRepositoryInterface) AdminServiceInterface {
	return &adminService{
		adminRepository: adminRepository,
	}
}

func (admin *adminService) CreateAdvisor(data request.CreateAdvisor) error {
	errLength := validation.CheckLength(data.Password)
	if errLength != nil {
		return errLength
	}

	role, errRole := validation.CheckEqual(data.Role, constanta.RoleType)
	if errRole != nil {
		return errRole
	}

	data.Role = role
	err := admin.adminRepository.CreateAdvisor(data)
	if err != nil {
		return err
	}

	return nil
}

func (admin *adminService) Login(data request.AdminLogin) (response.ResponseLogin, error) {
	dataAdmin, err := admin.adminRepository.FindNip(data.Id)
	if err != nil {
		return response.ResponseLogin{}, err
	}

	if !helper.CompareHash(dataAdmin.Password, data.Password) {
		return response.ResponseLogin{}, errors.New("error : password is wrong")
	}

	token, err := auth.CreateToken(dataAdmin.Id, dataAdmin.Role)
	if err != nil {
		return response.ResponseLogin{}, err
	}

	response := response.ModelToResponseLogin(dataAdmin.Name, dataAdmin.Role, token)
	return response, nil
}

func (admin *adminService) InsertList(data request.ListCollege) error {
	err := admin.adminRepository.InsertList(data)
	if err != nil {
		return err
	}

	return nil
}