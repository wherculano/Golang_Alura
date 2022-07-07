package main

import (
	"Alura/Go_Gin_Rest_API/database"
	"Alura/Go_Gin_Rest_API/routes"
)

func main() {
	database.ConnectDB()

	routes.HandlerRequests()
}
