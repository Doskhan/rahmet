package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rahmet/controllers"
	"rahmet/database"
)

func main() {

	// Load Configurations from config.json using Viper
	LoadAppConfig()

	// Initialize Database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()
	//database.Load(DB)

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterRoutes(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func RegisterRoutes(router *mux.Router) {
	// user
	router.HandleFunc("/api/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/user/{id}", controllers.GetUserById).Methods("GET")

	router.HandleFunc("/api/user", controllers.CreateUser).Methods("POST")

	// event
	router.HandleFunc("/api/events", controllers.GetAllEvents).Methods("GET")
	router.HandleFunc("/api/events/user/{id}", controllers.GetEventsByUserId).Methods("GET")

	router.HandleFunc("/api/event", controllers.CreateEvent).Methods("POST")
	router.HandleFunc("/api/event/participate", controllers.ParticipateEvent).Methods("POST")
	router.HandleFunc("/api/event/rahmet", controllers.RahmetEvent).Methods("POST")
	//router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	//router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	//router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")
}
