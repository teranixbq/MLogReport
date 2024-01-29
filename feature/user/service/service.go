package service

import (
	"errors"
	"mlogreport/feature/user/dto/request"
	"mlogreport/feature/user/repository"
	"mlogreport/utils/helper"
)

type userService struct {
	userRepository repository.UserRepositoryInterface
}

type UserServiceInterface interface {
	InsertUser(data request.RequestUser) error
	Login(data request.RequestLogin) error
}

func NewUserService(userRepository repository.UserRepositoryInterface) UserServiceInterface {
	return &userService{
		userRepository: userRepository,
	}
}

func (user *userService) InsertUser(data request.RequestUser) error {
	password,err := helper.HashPass(data.Password)
	if err !=nil {
		return err
	}

	data.Password = password
	err = user.userRepository.InsertUser(data)
	if err != nil {
		return err
	}

	return nil
}

func (user *userService) Login(data request.RequestLogin) error {
	dataUser,err := user.userRepository.FindNim(data.Nim)
	if err != nil {
		return err
	}

	if !helper.CompareHash(data.Password, dataUser.Password) {
		return errors.New("error : password salah")
	}

	return nil
}