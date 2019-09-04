package controllers

import (
	"log"
	"net/http"

	"github.com/jhindulak/go-rest-api-example/models"
)

var HealthCheck = func(w http.ResponseWriter, r *http.Request) {
	db := models.GetDB()

	if err := db.DB().Ping(); err != nil {
		log.Printf("Error connecting to database: %v", err)
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
