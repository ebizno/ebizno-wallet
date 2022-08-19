package postgres

import (
	"github.com/Paulo-Lopes-Estevao/e-bizno/adapter/wellet/infrastructure/database"
	"github.com/jinzhu/gorm"
	"os"
)

func ConnectedInPostgres() *gorm.DB {

	dsn := &database.ParamsDsn{
		Postgres: struct {
			DBHOST     string
			DBUSERNAME string
			DBPASSWORD string
			DBNAME     string
			DBPORT     string
		}{DBHOST: os.Getenv("DB_HOST"), DBUSERNAME: os.Getenv("DB_USERNAME"), DBPASSWORD: os.Getenv("DB_PASSWORD"), DBNAME: os.Getenv("DB_NAME"), DBPORT: os.Getenv("DB_PORT")},
	}

	return database.ConnectionDB("postgres", dsn)
}
