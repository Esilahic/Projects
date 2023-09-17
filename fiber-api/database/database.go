package database

import (
	"log"
	"os"

	"github.com/Esilahic/fiber-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseInstance struct {
	Db *gorm.DB
}

var Database DatabaseInstance

func Connect() {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database\n", err)
		os.Exit(1)
	}
	log.Println("Database connected")
	db.Logger = db.Logger.LogMode(logger.Info)
	log.Println("Running migrations...")
	// add migrations
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DatabaseInstance{Db: db}
}
