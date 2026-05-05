package config

import (
	"fmt"
	"os"

	"github.com/Mobilizes/materi-be-alpro/database/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("gagal konek ke database: " + err.Error())
	}

	// GORM auto-migrate — semua tabel dibuat otomatis dari struct
	db.AutoMigrate(
		&entities.User{},
		&entities.UserProfile{},
		&entities.MealPlan{},
		&entities.MealPlanVersion{},
		&entities.ChatHistory{},
	)

	DB = db
	return db
}
