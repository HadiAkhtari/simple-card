package main

import (
	"fmt"
	"keycard_service/initializer"
	"keycard_service/internal/database"
	"log"
)

func init() {
	initializer.ConnetToDB()
}
func main() {
	err := initializer.DB.AutoMigrate(
		&database.User{},
		&database.Terminal{},
		&database.KeyCard{},
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Database initialized")
}
