package config

import (
	"github.com/isaaclyra132/petshop-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:15433/postgres"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{}, &models.Pet{}, &models.Vaccine{}, &models.Veterinarian{}, &models.Appointment{})

	DB = db
}
