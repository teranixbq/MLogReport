package database

import (
	"fmt"
	"mlogreport/app/config"
	user "mlogreport/feature/user/model"
	admin "mlogreport/feature/admin/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDBPostgres(cfg *config.Config) *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHOST, cfg.DBUSER, cfg.DBPASS, cfg.DBNAME, cfg.DBPORT)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true, 
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.Users{})
	db.AutoMigrate(&admin.Admin{})
}