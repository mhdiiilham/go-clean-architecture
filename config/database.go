package config

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var db *gorm.DB
var port string

func init() {
	log.Info("Function ini called")
	if err := godotenv.Load(); err != nil {
		log.Fatal("Cannot read .env file")
	}
}

func SetupDB() {
	log.Info("Setting up DB")

	host := os.Getenv("DB_HOST")
	db := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		db,
	)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database.", err)
	}

	SetUpDBConnection(conn)
	SetPortConnection(os.Getenv("PORT"))
}

func SetUpDBConnection(DB *gorm.DB) {
	db = DB
}
func GetDBConnection() *gorm.DB {
	return db
}

func SetPortConnection(Port string) {
	port = Port
}

func GetPortConnection() string {
	return port
}
