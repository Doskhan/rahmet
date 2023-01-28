package main

import (
	"fmt"
	"log"
	"net/http"
	"rahmet/controllers"
	"rahmet/database"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	// Load Configurations from config.json using Viper
	LoadAppConfig()

	// Initialize Database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()
	database.Load(DB)

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterRoutes(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/users", controllers.GetAllUsers).Methods("GET")
	//router.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("GET")
	//router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	//router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	//router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")
}
