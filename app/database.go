package app

import (
	"github.com/gowesmart/api-gowesmart/model/entity"
	"github.com/gowesmart/api-gowesmart/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(utils.MustGetEnv("DB_DSN")), &gorm.Config{
		QueryFields: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})
	utils.PanicIfError(err)

	err = db.AutoMigrate(&entity.User{}, &entity.Profile{}, &entity.Role{})
	utils.PanicIfError(err)

	return db
}
