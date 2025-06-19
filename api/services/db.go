package services

import (
	"os"

	"github.com/riju-stone/sevin/api/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	l := utils.CustomLogger
	dbConn := os.Getenv("DB_CONN")

	conn, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
	if err != nil {
		l.Fatalf("Failed to connect to the database: %v", err)
	}
	DB = conn
	l.Info("Connected to the database")
}
