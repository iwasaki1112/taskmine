package database

import (
	"taskmine/config"

	"gorm.io/gorm"
)

type DbConnector interface {
	Connect(config *config.Config) error
	AutoMigration(models ...interface{}) error
	GetDB() *gorm.DB
}
