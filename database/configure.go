package database

import (
	"fmt"
	"os"

	"github.com/jhindulak/go-rest-api-example/models"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// OpenDB opens connection do DB with credentials
func OpenDB() *gorm.DB {
	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}

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

	db, err := gorm.Open("postgres", dbUri)
	if err != nil {
		panic(err)
	}

	if err := db.DB().Ping(); err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(&models.Account{}, &models.Contact{})

	return db
}

// SetupDB runs migrations and seeds (if flag is set up)
func SetupDB() {
	db := OpenDB()

	if seedDatabase := os.Getenv("RUN_SEEDS"); seedDatabase == "true" {
		runSeeds(db)
	}
}
