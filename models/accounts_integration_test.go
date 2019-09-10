// +build integration

package models

import (
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

// Ensure the database connection is made so functions under test can access the DB
func init() {
	RegisterTxDB("txdb")
}

func TestStoreType_CreateAccount_ValidAccount(t *testing.T) {
	db, _ := PrepareTestDB("txdb")
	defer CleanTestDB(db)

	store := &StoreType{DB: db}

	type TestCase struct {
		name            string
		account         *Account
		expectedStatus  bool
		expectedMessage string
	}

	var testCases []TestCase

	for i := 0; i < 10; i++ {
		testCases = append(testCases, TestCase{
			name: faker.Internet().Email(),
			account: &Account{
				Email:    faker.Internet().Email(),
				Password: faker.Internet().Password(6, 20),
				Token:    faker.RandomString(20),
			},
			expectedStatus:  true,
			expectedMessage: "Account created",
		})
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := store.CreateAccount(tc.account)
			acc := &Account{}
			store.DB.Table("accounts").Where("email = ?", tc.account.Email).First(acc)

			require.Equal(t, tc.expectedStatus, result["status"])
			require.Equal(t, tc.expectedMessage, result["message"])
			require.Equal(t, tc.account.Email, acc.Email)
		})
	}
}

func TestStoreType_CreateAccount_AccountNotCreatedInvalidEmail(t *testing.T) {
	db, _ := PrepareTestDB("txdb")
	defer CleanTestDB(db)

	store := &StoreType{DB: db}

	account := &Account{
		Email:    "",
		Password: faker.Internet().Password(7, 20),
		Token:    faker.RandomString(20),
	}

	result := store.CreateAccount(account)

	acc := &Account{}
	err := store.DB.Table("accounts").Where("email = ?", account.Email).First(acc).Error

	require.Equal(t, false, result["status"])
	require.Equal(t, "Valid email address is required", result["message"])
	require.Equal(t, gorm.ErrRecordNotFound, err)
	require.Error(t, err)
}

func TestStoreType_CreateAccount_AccountNotCreatedInvalidPassword(t *testing.T) {
	db, _ := PrepareTestDB("txdb")
	defer CleanTestDB(db)

	store := &StoreType{DB: db}

	account := &Account{
		Email:    faker.Internet().Email(),
		Password: faker.Internet().Password(0, 5),
		Token:    faker.RandomString(20),
	}

	result := store.CreateAccount(account)

	acc := &Account{}
	err := store.DB.Table("accounts").Where("email = ?", account.Email).First(acc).Error

	require.Equal(t, false, result["status"])
	require.Equal(t, "Password is required and must be longer than 6 characters", result["message"])
	require.Equal(t, gorm.ErrRecordNotFound, err)
	require.Error(t, err)
}

func TestStoreType_CreateAccount_AccountNotCreatedIfExists(t *testing.T) {
	db, _ := PrepareTestDB("txdb")
	defer CleanTestDB(db)

	store := &StoreType{DB: db}

	account1 := &Account{
		Email:    "thisemailshouldalreadyexist@email.com",
		Password: faker.Internet().Password(7, 20),
		Token:    faker.RandomString(20),
	}

	account2 := &Account{
		Email:    "thisemailshouldalreadyexist@email.com",
		Password: faker.Internet().Password(7, 20),
		Token:    faker.RandomString(20),
	}

	store.CreateAccount(account1)
	duplicateAccountResult := store.CreateAccount(account2)

	var accountCount int

	store.DB.Table("accounts").Where("email = ?", account1.Email).Count(&accountCount)

	require.Equal(t, false, duplicateAccountResult["status"])
	require.Equal(t, "An account with this email address already exists", duplicateAccountResult["message"])
	require.Equal(t, 1, accountCount)
}
