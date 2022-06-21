package main

import (
	"html/template"
	"net/http"
)


var temp8 = template.Must(template.ParseGlob("templates/*.html"))

func main(){
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	temp8.ExecuteTemplate(w, "Index", nil)
}