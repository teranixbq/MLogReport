package service

import (
	"mlogreport/feature/user/repository"
	"mlogreport/feature/user/dto/request"
	"mlogreport/utils/helper"
)

type userService struct {
	userRepository repository.UserRepositoryInterface
}

type UserServiceInterface interface {
	InsertUser(data request.RequestUser) error
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
