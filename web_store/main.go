package main

import (
	"html/template"
	"net/http"
)


type Product struct{
	Name string
	Description string
	Value float64
	Amount int
}


var temp8 = template.Must(template.ParseGlob("templates/*.html"))

func main(){
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	products := []Product{
		{Name: "Earphone", Description: "New!", Value: 12.34, Amount: 6},
		{"Flash Drive", "32Gb", 9.99, 10},
		{"Keyboard", "With leds", 20, 4},
	}
	
	temp8.ExecuteTemplate(w, "Index", products)
}