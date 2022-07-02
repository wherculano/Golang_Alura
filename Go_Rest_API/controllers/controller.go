package controllers

import (
	"Go_Rest_API/database"
	"Go_Rest_API/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality
	database.DB.First(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var newPersonality models.Personality
	json.NewDecoder(r.Body).Decode(&newPersonality)
	database.DB.Create(&newPersonality)
	json.NewEncoder(w).Encode(newPersonality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality
	database.DB.Delete(&personality, id)
	json.NewEncoder(w).Encode(personality)
}
