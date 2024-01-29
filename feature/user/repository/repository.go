package repository

import (
	"mlogreport/feature/user/dto/request"
	"mlogreport/feature/user/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	InsertUser(data request.RequestUser) error
	FindNim(nim string) (model.Users,error)
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

func (user *userRepository) FindNim(nim string) (model.Users,error) {
	dataUser := model.Users{}

	tx := user.db.Where("nim = ?", nim).First(&dataUser)
	if tx.Error != nil {
		return dataUser,tx.Error
	}

	return dataUser,nil
}

