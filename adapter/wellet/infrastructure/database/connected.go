package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ParamsDsn struct {
	Postgres struct {
		DBHOST     string
		DBUSERNAME string
		DBPASSWORD string
		DBNAME     string
		DBPORT     string
	}
}

func ConnectionDB(drive string, Indsn *ParamsDsn) *gorm.DB {

	var dsn string
	if drive == "postgres" {
		dsn = "host=" + Indsn.Postgres.DBHOST + " user=" + Indsn.Postgres.DBUSERNAME + " password=" + Indsn.Postgres.DBPASSWORD + " dbname=" + Indsn.Postgres.DBNAME + " port=" + Indsn.Postgres.DBPORT + " sslmode=disable"
	}

	db, err := gorm.Open(drive, dsn)

	if err != nil {

		panic("failed to connect database")

	}

	fmt.Println("Connection Opened to Database")

	return db

}
