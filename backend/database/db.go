package database

import (
	"backend/database/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("ebiznes.db"))
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Payment{})
	db.AutoMigrate(&models.Order{})

	return db
}
