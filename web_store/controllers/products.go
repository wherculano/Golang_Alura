package controllers

import (
	"Alura/web_store/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp8 = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	all_products := models.GetAllProducts()

	temp8.ExecuteTemplate(w, "Index", all_products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp8.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error during parse price:", err)
		}

		amountInt, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Error during parse amount:", err)
		}

		models.CreateNewProduct(name, description, priceFloat, amountInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	produtId := r.URL.Query().Get("id")
	models.DeleteProductById(produtId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	log.Print("ID recebido", productId)
	product := models.EditProduct(productId)

	temp8.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		idToInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error during parsing ID")
		}

		amountToInt, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Error during parsing Amount")
		}

		priceToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error during parsing Price")
		}

		models.UpdateProductFields(idToInt, name, description, priceToFloat, amountToInt)
	}
	http.Redirect(w, r, "/", 301)
}
