package controllers

import (
	"Alura/web_store/models"
	"net/http"
	"text/template"
)

var temp8 = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	all_products := models.GetAllProducts()

	temp8.ExecuteTemplate(w, "Index", all_products)
}
