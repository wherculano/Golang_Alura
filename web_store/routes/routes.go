package routes

import (
	"Alura/web_store/controllers"
	"net/http"
)

func GetRoutes() {
	http.HandleFunc("/", controllers.Index)
}
