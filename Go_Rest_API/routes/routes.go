package routes

import (
	"Go_Rest_API/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalities", controllers.GetAllPersonalities)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
