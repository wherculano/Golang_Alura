package main

import (
	"Alura/Go_Gin_Rest_API/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func SetupTestRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestStatusCodeGreeting(t *testing.T) {
	r := SetupTestRoutes()
	r.GET("/:name", controllers.Greeting)
	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Fatalf("Status Code %d is different than %d.", resp.Code, http.StatusOK)
	}
}
