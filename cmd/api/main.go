package main

import (
	"gin-microservices-skeleton/internal/api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Set up middleware here if needed

	routes.SetupRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
