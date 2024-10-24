package main

import (
	"fmt"
	"log"
	"os"

	handlers "github.com/caciano/portfolio/internal/handlers"
	"github.com/caciano/portfolio/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	loadDotEnv()
	db := initDB()
	migrateDB(db)
	loadHandlers(db)
}

// Load all the routes and start the server
func loadHandlers(db *gorm.DB) {
	r := gin.Default()
	handlers.InitRoutes(r, db)
	r.Run(":80")
}

// Load the environment variables
func loadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("Error to load .env file")
	}
}

// Migrate the database
func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Project{}, &models.PersonalData{}, &models.Address{})
}

// Initialize the database
func initDB() *gorm.DB {
	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")
	port := os.Getenv("DATABASE_PORT")
	sslmode := os.Getenv("DATABASE_SSLMODE")
	timezone := os.Getenv("DATABASE_TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslmode, timezone)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
		panic("failed to connect database")
	}
	log.Println("connected to database")
	return db
}
