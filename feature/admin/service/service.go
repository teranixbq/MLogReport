package service

import (
	"errors"
	"fmt"

	"mlogreport/feature/admin/dto/request"
	"mlogreport/feature/admin/dto/response"
	"mlogreport/feature/admin/repository"
	"mlogreport/utils/auth"
	"mlogreport/utils/constanta"
	"mlogreport/utils/enum"
	"mlogreport/utils/helper"
	"mlogreport/utils/meta"
	"mlogreport/utils/validation"
)

type adminService struct {
	adminRepository repository.AdminRepositoryInterface
}

type AdminServiceInterface interface {
	CreateAdvisor(data request.CreateAdvisor) error
	Login(data request.AdminLogin) (response.ResponseLogin, error)
	SelectAllAdvisor(page, limit int) ([]response.ResponseAllAdvisor, meta.Meta, error)
	SelectAdvisor(id string) (response.ResponseAdvisor, error)
	InsertList(data request.ListCollege) error
	DeleteAdvisor(id string) error
}

func NewAdminService(adminRepository repository.AdminRepositoryInterface) AdminServiceInterface {
	return &adminService{
		adminRepository: adminRepository,
	}
}

func (admin *adminService) CreateAdvisor(data request.CreateAdvisor) error {
	errEmpty := validation.CheckEmpty(data.Nip, data.Name, data.Password, data.Role)
	if errEmpty != nil {
		return errEmpty
	}

	errLength := validation.CheckLength(data.Password)
	if errLength != nil {
		return errLength
	}

	role, errRole := validation.CheckEqual(data.Role, enum.RoleType)
	if errRole != nil {
		return errRole
	}

	dataUser, _ := admin.adminRepository.SelectNip(data.Nip)
	if dataUser.Nip != "" {
		dataErr := fmt.Sprintf(constanta.EXISTS, dataUser.Nip)
		return errors.New(dataErr)
	}

	password, errHash := helper.HashPass(data.Password)
	if errHash != nil {
		return errHash
	}

	data.Password = password
	data.Role = role
	err := admin.adminRepository.CreateAdvisor(data)
	if err != nil {
		return err
	}

	return nil
}

func (admin *adminService) Login(data request.AdminLogin) (response.ResponseLogin, error) {
	errEmpty := validation.CheckEmpty(data.Nip, data.Password)
	if errEmpty != nil {
		return response.ResponseLogin{}, errEmpty
	}

	dataAdmin, err := admin.adminRepository.SelectNip(data.Nip)
	if err != nil {
		return response.ResponseLogin{}, err
	}

	if !helper.CompareHash(dataAdmin.Password, data.Password) {
		return response.ResponseLogin{}, errors.New(constanta.WRONG_PASS)
	}

	token, err := auth.CreateToken(dataAdmin.Id, dataAdmin.Role)
	if err != nil {
		return response.ResponseLogin{}, err
	}

	response := response.ModelToResponseLogin(dataAdmin.Name, dataAdmin.Role, token)
	return response, nil
}

func (admin *adminService) SelectAllAdvisor(page, limit int) ([]response.ResponseAllAdvisor, meta.Meta, error) {
	page, limit, err := validation.CheckPagination(page, limit)
	if err != nil {
		return nil, meta.Meta{}, err
	}

	dataAdvisor, metaInfo, err := admin.adminRepository.SelectAllAdvisor(page, limit)
	if err != nil {
		return nil, meta.Meta{}, err
	}

	if metaInfo.TotalPage != 0 {
		if page > metaInfo.TotalPage {
			return nil, meta.Meta{}, errors.New("error : page is not available")
		}
	}

	return dataAdvisor, metaInfo, nil
}

func (admin *adminService) SelectAdvisor(id string) (response.ResponseAdvisor, error) {
	dataAdmin, err := admin.adminRepository.SelectAdvisor(id)
	if err != nil {
		return response.ResponseAdvisor{}, err
	}
	return dataAdmin, nil
}

func (admin *adminService) InsertList(data request.ListCollege) error {
	errEmpty := validation.CheckEmpty(data.Advisor,data.Colleges)
	if errEmpty != nil {
		return errEmpty
	}
	err := admin.adminRepository.InsertList(data)
	if err != nil {
		return err
	}

	return nil
}

func (admin *adminService) DeleteAdvisor(id string) error {
	dataAdvisor, err := admin.adminRepository.SelectAdvisor(id)
	if err != nil {
		return err
	}

	if len(dataAdvisor.Colleges) != 0 {
		return errors.New("error: Cannot delete advisors currently connected to the college")
	}

	err = admin.adminRepository.DeleteAdvisor(id)
	if err != nil {
		return err
	}

	return nil
}
