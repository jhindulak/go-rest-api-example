package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jhindulak/go-rest-api-example/models"
	"github.com/jhindulak/go-rest-api-example/utils"
)

func (local StoreType) CreateAccount(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	fmt.Printf("Creating account for email: %v", account.Email)
	resp := local.Store.CreateAccount(account)
	utils.Respond(w, resp)
}

func (local StoreType) Authenticate(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := local.Store.Login(account.Email, account.Password)
	utils.Respond(w, resp)
}
