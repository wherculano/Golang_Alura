package main

import (
	"Alura/Go_Gin_Rest_API/controllers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupTestRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestStatusCodeGreeting(t *testing.T) {
	r := SetupTestRoutes()
	r.GET("/:name", controllers.Greeting)
	req, _ := http.NewRequest("GET", "/Wagner", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Status Code %d is different than %d.", resp.Code, http.StatusOK)

	mockResponse := `{"message":"Hello Wagner"}`
	responseBody, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, mockResponse, string(responseBody))

}
