package app

import (
	"github.com/gowesmart/api-gowesmart/helper"
	"github.com/gowesmart/api-gowesmart/model/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(helper.MustGetEnv("DB_DSN")), &gorm.Config{
		QueryFields: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})
	helper.PanicIfError(err)

	err = db.AutoMigrate(&entity.User{})
	helper.PanicIfError(err)

	return db
}
