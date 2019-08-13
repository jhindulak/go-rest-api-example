package models

import (
	"fmt"

	"github.com/jhindulak/go-rest-api-example/utils"
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserId uint   `json:"user_id"`
}

func (contact *Contact) Validate() (map[string]interface{}, bool) {
	if contact.Name == "" {
		return utils.Message(false, "Contact name cannot be empty"), false
	}

	if contact.Phone == "" {
		return utils.Message(false, "Phone number cannot be empty"), false
	}

	if contact.UserId <= 0 {
		return utils.Message(false, "User ID is not recognized"), false
	}

	return utils.Message(true, "Successfully validated contact"), true
}

func (contact *Contact) Create() map[string]interface{} {
	if resp, ok := contact.Validate(); !ok {
		return resp
	}

	GetDB().Create(contact)

	resp := utils.Message(true, "Successfully created contact")
	resp["contact"] = contact

	return resp
}

func GetContact(id uint) *Contact {
	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}

	return contact
}

func GetContacts(user uint) []*Contact {
	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contacts
}
