package database

import (
	"log"
	"os"

	"github.com/storyofhis/microservice-go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var Database DBInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	
	if err != nil {
		log.Fatal("Failed connected to Database! \n", err.Error())
		os.Exit(2)
	}
	log.Println("Connected to Database successfully!")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migration...")

	// migration
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
	
	Database = DBInstance{Db: db}
} 