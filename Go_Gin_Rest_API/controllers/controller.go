package controllers

import (
	"Alura/Go_Gin_Rest_API/models"

	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	c.JSON(200, models.Students)
}

func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"message": "Hello " + name,
	})
}
