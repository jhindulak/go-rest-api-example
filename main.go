package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jhindulak/go-rest-api-example/controllers"

	"github.com/gorilla/mux"
	"github.com/jhindulak/go-rest-api-example/app"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) // Attach middleware JWT auth

	// Authentication Handlers
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	// Contact Handlers
	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET")

	port := os.Getenv("listen_port")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("Listening on port: %q", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
