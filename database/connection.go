package database

import (
	"github.com/cestevezing/eventos/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Connect() (*gorm.DB, error) {
	conf := config.GetDatabaseConfig()
	dsn := "host=" + conf.Host +
		" port=" + conf.Port +
		" user=" + conf.User +
		" dbname=" + conf.Name +
		" password=" + conf.Password +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db
	return db, nil
}
