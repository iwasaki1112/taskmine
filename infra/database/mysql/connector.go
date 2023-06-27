package mysql

import (
	"fmt"
	"taskmine/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connector struct {
	DB *gorm.DB
}

func NewConnector() *Connector {
	return &Connector{}
}

func (m *Connector) GetDB() *gorm.DB {
	return m.DB
}

func (m *Connector) Connect(dbConfig *config.Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBName)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	m.DB = DB
	return nil
}

func (m *Connector) AutoMigration(models ...interface{}) error {
	err := m.DB.AutoMigrate(models)
	if err != nil {
		return err
	}
	return nil
}
