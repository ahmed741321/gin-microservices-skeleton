package test

import (
	"gin-microservices-skeleton/internal/api/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetExample(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Set up the controller and routes
	exampleController := controllers.ExampleController{}
	router.GET("/example", exampleController.GetExample)

	// Create a request to send to the above route
	req, _ := http.NewRequest(http.MethodGet, "/example", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "example response"}`, w.Body.String())
}

func TestCreateExample(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Set up the controller and routes
	exampleController := controllers.ExampleController{}
	router.POST("/example", exampleController.CreateExample)

	// Create a request to send to the above route
	req, _ := http.NewRequest(http.MethodPost, "/example", nil) // Add body as needed
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, `{"message": "example created"}`, w.Body.String())
}
