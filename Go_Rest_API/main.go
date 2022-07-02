package main

import (
	"Go_Rest_API/database"
	"Go_Rest_API/models"
	"Go_Rest_API/routes"
	"fmt"
)

func main() {
	fmt.Println("Starting Server ...")

	models.Personalities = []models.Personality{
		{Id: 1, Name: "Name 1", History: "History 1"},
		{Id: 2, Name: "Name 2", History: "History 2"},
	}

	database.ConnectDB()
	routes.HandleRequest()
}
