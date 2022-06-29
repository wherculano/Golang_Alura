package main

import (
	"Alura/web_store/routes"
	"net/http"
)

func main() {
	routes.GetRoutes()
	http.ListenAndServe(":8000", nil)
}
