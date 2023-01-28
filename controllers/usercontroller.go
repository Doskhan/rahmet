package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"rahmet/database"
	"rahmet/models"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.Instance.Order("bonus desc").Find(&users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]
	if checkIfUserExists(userId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var user models.User
	database.Instance.First(&user, userId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	database.Instance.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func checkIfUserExists(userId string) bool {
	var product models.User
	database.Instance.First(&product, userId)
	if product.ID == 0 {
		return false
	}
	return true
}
