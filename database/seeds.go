package database

import (
	"math/rand"
	"time"

	"github.com/jhindulak/go-rest-api-example/models"
	"github.com/jinzhu/gorm"
	"syreclabs.com/go/faker"
)

func runSeeds(db *gorm.DB) {
	seedAccounts(db)
	seedContacts(db)
}

func seedAccounts(db *gorm.DB) {
	for count := 0; count < 100; count++ {
		db.Create(&models.Account{
			Email:    faker.Internet().Email(),
			Password: faker.Internet().Password(6, 20),
		})
	}
}

func seedContacts(db *gorm.DB) {
	for count := uint(0); count < 100; count++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		randomContactCount := r.Intn(0-10) + 0
		for count2 := 0; count2 < randomContactCount; count2++ {
			db.Create(&models.Contact{
				Name:   faker.Name().Name(),
				Phone:  faker.PhoneNumber().PhoneNumber(),
				UserId: count,
			})
		}
	}
}
