package services

import (
	"os"

	"github.com/riju-stone/sevin/api/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	l := utils.CustomLogger
	dbConn := os.Getenv("DB_CONN")
	l.Debugf("Connecting to the database: %s", dbConn)

	conn, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	l.Info("Connected to the database")
	return conn, nil
}
