package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jhindulak/go-rest-api-example/database"
	"github.com/jhindulak/go-rest-api-example/models"

	"github.com/jhindulak/go-rest-api-example/controllers"

	"github.com/gorilla/mux"
	"github.com/jhindulak/go-rest-api-example/app"
)

func main() {
	db := database.OpenDB()
	store := &models.StoreType{DB: db}

	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) // Attach middleware JWT auth

	// Authentication Handlers
	router.HandleFunc("/api/user/new", controllers.StoreType{Store: store}.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.StoreType{Store: store}.Authenticate).Methods("POST")

	// Contact Handlers
	router.HandleFunc("/api/me/contacts", controllers.StoreType{Store: store}.GetContactsFor).Methods("GET")

	// Health Check Handler
	router.HandleFunc("/api/healthcheck", controllers.StoreType{Store: store}.HealthCheck).Methods("GET")

	fmt.Println("Finished adding handlers...")

	port := os.Getenv("listen_port")
	if port == "" {
		port = "80"
	}

	fmt.Println("Listening on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
