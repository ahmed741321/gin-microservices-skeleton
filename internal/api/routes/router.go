package routes

import (
	"gin-microservices-skeleton/internal/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/example", controllers.NewExampleController().CreateExample)
		api.POST("/example", controllers.NewExampleController().GetExample)
	}

	return router
}
