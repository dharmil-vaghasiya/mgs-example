package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

var DB_URI = "mongodb://localhost:27017" // for running it in local
const DB_NAME = "mgs-example"

func main() {
	log.Print(fmt.Sprintf("Running service in %v environment", os.Getenv("Environment")))

	if os.Getenv("Environment") == "Production" {
		DB_URI = os.Getenv("DB_URI")
		gin.SetMode(gin.ReleaseMode)
	}

	// dependency injection
	repo := NewRepository(DB_URI)
	repo.LoadData()
	service := NewService(repo)

	// registering routes
	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/products", service.GetProducts)

	// Run server
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatal("unable to start the server")
		return
	}
}

// Ping route
func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
