package initializer

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnetToDB() {
	var err error
	connStr := "host=localhost port=5432 user=postgres password=1234 dbname=keycard sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
}
