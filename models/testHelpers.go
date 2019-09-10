package models

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/jinzhu/gorm"

	// needs to be here because of the DB Open
	_ "github.com/lib/pq"
)

func RegisterTxDB(name string) {
	username := os.Getenv("MASTER_USERNAME")
	password := os.Getenv("MASTER_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("ENDPOINT_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
		fmt.Println("Using default DB port 5432.")
	}

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		dbHost, dbPort, username, dbName, password)

	txdb.Register(name, "postgres", dbUri)
}

func PrepareTestDB(withName string) (*gorm.DB, error) {
	sqlDB, err := sql.Open(withName, fmt.Sprintf("connection_%d", time.Now().UnixNano()))
	db, err := gorm.Open("postgres", sqlDB)

	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(&Account{}, &Contact{})

	return db, err
}

func CleanTestDB(db *gorm.DB) {
	_ = db.Close()
}
