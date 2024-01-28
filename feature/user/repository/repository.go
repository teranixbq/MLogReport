package repository

import (
	"mlogreport/feature/user/dto/request"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	InsertUser(data request.RequestUser) error
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}

func (user *userRepository) InsertUser(data request.RequestUser) error {
	input := request.RequestUserToModel(data)

	tx := user.db.Create(&input)
	if tx.Error != nil {
		return nil
	}
	
	return nil
}