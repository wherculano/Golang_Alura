package main

import(
	"github.com/gin-gonic/gin"
)


func GetAllStudents(c *gin.Context){
	c.JSON(200, gin.H{
		"id": "1",
		"name": "Wagner Herculano",
	})
}

func main(){
	req := gin.Default()
	req.GET("/students", GetAllStudents)
	req.Run()
}

