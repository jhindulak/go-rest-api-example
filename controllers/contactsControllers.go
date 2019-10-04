package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/jhindulak/go-rest-api-example/models"
	"github.com/jhindulak/go-rest-api-example/utils"
	"net/http"
)

func (local StoreType) CreateContact(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserId = user
	resp := local.Store.CreateContact(contact)
	utils.Respond(w, resp)
}

func (local StoreType) GetContactsFor(w http.ResponseWriter, r *http.Request) {
	id, found := r.Context().Value("user").(uint)

	if !found {
		utils.Respond(w, utils.Message(false, "There was an error in your request"))
		return
	}

	data := local.Store.GetContacts(id)
	fmt.Printf("Getting contacts for userId: %d", id)
	resp := utils.Message(true, "Successfully retrieved contacts")
	resp["data"] = data
	utils.Respond(w, resp)
}
