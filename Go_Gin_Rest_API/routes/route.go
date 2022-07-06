package routes

import (
	"Alura/Go_Gin_Rest_API/controllers"

	"github.com/gin-gonic/gin"
)

func HandlerRequests() {
	req := gin.Default()
	req.GET("/students", controllers.GetAllStudents)
	req.GET("/:name", controllers.Greeting)
	req.Run()
}
