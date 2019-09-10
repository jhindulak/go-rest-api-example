package controllers

import (
	"log"
	"net/http"
)

func (local StoreType) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if err := local.Store.DB.DB().Ping(); err != nil {
		log.Printf("Error connecting to database: %v", err)
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
