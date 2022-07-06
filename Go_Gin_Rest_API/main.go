package main

import (
	"Alura/Go_Gin_Rest_API/models"
	"Alura/Go_Gin_Rest_API/routes"
)

func main() {
	models.Students = []models.Student{
		{Name: "Wagner Herculano", CPF: "111.222.333-45", RG: "01.234.567-8"},
		{Name: "Herculano Wagner", CPF: "222.333.444-56", RG: "12.345.678-9"},
	}
	routes.HandlerRequests()
}
