package database

import (
	"api-avaliacao/src/config"
	"api-avaliacao/src/models"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(postgres.Open(config.StrConnectDB), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db = database

	config, _ := db.DB()

	config.SetConnMaxIdleTime(10)
	config.SetConnMaxLifetime(time.Hour)

	err = db.AutoMigrate(models.Pessoa{})
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
