package routes

import (
	"Go_Rest_API/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	new_router := mux.NewRouter()
	new_router.HandleFunc("/", controllers.Home)
	new_router.HandleFunc("/api/personalities", controllers.GetAllPersonalities)
	log.Fatal(http.ListenAndServe(":8000", new_router))
}
