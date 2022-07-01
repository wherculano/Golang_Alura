package controllers

import (
	"Go_Rest_API/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(models.Personalities)
}