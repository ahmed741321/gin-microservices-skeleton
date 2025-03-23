package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExampleController struct {
	// Add any dependencies here, such as services
}

func NewExampleController() *ExampleController {
	return &ExampleController{}
}

func (ec *ExampleController) GetExample(c *gin.Context) {
	// Logic to get an example
	c.JSON(http.StatusOK, gin.H{"message": "Get Example"})
}

func (ec *ExampleController) CreateExample(c *gin.Context) {
	// Logic to create an example
	c.JSON(http.StatusCreated, gin.H{"message": "Example Created"})
}
