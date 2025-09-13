package main

import (
	"github.com/gin-gonic/gin"
	"keycard_service/internal/database"

	"keycard_service/initializer"

	"log"
)

func init() {
	initializer.ConnetToDB()
}

func main() {
	models := database.NewModels(initializer.DB)
	g := gin.Default()
	v1 := g.Group("api/v1")
	keyHandler := NewKeyCardHandler(models.KeyCards)
	v1.POST("/keyservice", keyHandler.CreatePOSData)
	log.Println("Listening on port 8080")
	g.Run(":8080")

}
