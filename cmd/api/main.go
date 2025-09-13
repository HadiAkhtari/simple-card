package main
import (
	"github.com/gin-gonic/gin"

	"keycard_service/initializer"
	"log"

)

func init() {
	initializer.ConnetToDB()
}

func main() {
	g := gin.Default()
	v1 := g.Group("api/v1"){
		v1.POST("/keyservice",createKeyService)
		v1.GET("/keyservice",getAllKeyService)
	}
	log.Println("Listening on port 8080")
	g.Run(":8080")

}
