package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	err := godotenv.Load()

	if err != nil {
		log.Println("No .env found")
	}

	cs := GetConnectionString(false)
	db, err := gorm.Open(postgres.Open(cs), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed")
		log.Fatal("CS:", cs)
	}

	DB = db
}
func GetConnectionString(use_dsn bool) string {
	if use_dsn {
		var cfg Config
		file, _ := os.ReadFile("config.json")
		json.Unmarshal(file, &cfg)
		cs1 := cfg.ConnectionStrings.DefaultConnection
		return cs1
	} else {
		cs2 := "host=localhost user=postgres password=Iamsmart27! dbname=northwind port=5432 sslmode=disable"
		return cs2
	}
}

type Config struct {
	ConnectionStrings struct {
		DefaultConnection string
	}
}
