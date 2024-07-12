package database

import (
	"fmt"

	"github.com/mpt-tiendc/golang-template/internal/config"
	"github.com/mpt-tiendc/golang-template/internal/domain/entities"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func NewDatabase(dbconfig *config.DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(getDialector(dbconfig), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&entities.User{})
	return db, err
}

func getDialector(config *config.DatabaseConfig) gorm.Dialector {
	switch config.Type {
	case "mysql":
		return mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Username, config.Password, config.Host, config.Port, config.Database))
	case "postgres":
		return postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.Host, config.Username, config.Password, config.Database, config.Port))
	case "sqlite":
		return sqlite.Open(fmt.Sprintf("%s.db", config.Database))
	case "sqlserver":
		return sqlserver.Open(fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", config.Username, config.Password, config.Host, config.Port, config.Database))
	default:
		return nil
	}
}
