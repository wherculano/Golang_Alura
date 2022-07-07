package routes

import (
	"Alura/Go_Gin_Rest_API/controllers"

	"github.com/gin-gonic/gin"
)

func HandlerRequests() {
	req := gin.Default()
	req.GET("/students", controllers.GetAllStudents)
	req.GET("/:name", controllers.Greeting)
	req.POST("/students", controllers.CreateNewStudent)
	req.GET("/students/:id", controllers.GetStudentById)
	req.DELETE("/students/:id", controllers.DeleteStudent)
	req.PATCH("/students/:id", controllers.EditStudent)
	req.GET("/students/cpf/:cpf", controllers.GetStudentByCPF)
	req.Run()
}
