package lib

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	connectionString := "host=localhost port=5432 user=postgres password=marco123 dbname=assignment2 sslmode=disable"

	return gorm.Open(postgres.Open(connectionString), &gorm.Config{})
}
