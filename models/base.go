package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}

	username := os.Getenv("MASTER_USERNAME")
	password := os.Getenv("MASTER_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("ENDPOINT_ADDRESS")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		dbHost, username, dbName, password)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Println(err)
		log.Fatal("BROKEN")
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &Contact{})
}

func GetDB() *gorm.DB {
	return db
}
