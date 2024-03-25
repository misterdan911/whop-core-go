package app

import (
	// "fmt"
	"log"
	"whop-core-go/db"

	"whop-core-go/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var GinEngine *gin.Engine

func Init() {

	// load .env
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to database
	db.ConnectWhopDb();

	// Initiate gin
	GinEngine = gin.Default()

	// setup route
	routes.Setup(GinEngine)
}
