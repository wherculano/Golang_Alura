package main

import (
	"Alura/Go_Gin_Rest_API/controllers"
	"Alura/Go_Gin_Rest_API/database"
	"Alura/Go_Gin_Rest_API/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) // less info on terminal about the test
	routes := gin.Default()
	return routes
}

func CreateMockStudent() {
	student := models.Student{Name: "Name Test", CPF: "12345678900", RG: "98765432"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteMockStudent() {
	var student models.Student
	database.DB.Delete(&student, ID)
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

func TestGetAllStudents(t *testing.T) {
	database.ConnectDB()
	CreateMockStudent()
	defer DeleteMockStudent()

	r := SetupTestRoutes()
	r.GET("/students", controllers.GetAllStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetStudentByCPF(t *testing.T) {
	database.ConnectDB()
	CreateMockStudent()
	defer DeleteMockStudent()

	r := SetupTestRoutes()
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCPF)
	req, _ := http.NewRequest("GET", "/students/cpf/12345678900", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetStudentByID(t *testing.T) {
	database.ConnectDB()
	CreateMockStudent()
	defer DeleteMockStudent()

	r := SetupTestRoutes()
	r.GET("/students/:id", controllers.GetStudentById)

	endpoint := "/students/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("GET", endpoint, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	var studentMock models.Student
	json.Unmarshal(resp.Body.Bytes(), &studentMock)

	assert.Equal(t, "Name Test", studentMock.Name)
	assert.Equal(t, "12345678900", studentMock.CPF)
	assert.Equal(t, "98765432", studentMock.RG)
}

func TestDeleteStudent(t *testing.T) {
	database.ConnectDB()
	CreateMockStudent()

	r := SetupTestRoutes()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	endpoint := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", endpoint, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}
