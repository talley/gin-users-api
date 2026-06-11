package config

import (
	"encoding/json"
	"fmt"
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
	if use_dsn == true {
		var cfg Config
		file, _ := os.ReadFile("config.json")
		json.Unmarshal(file, &cfg)
		cs1 := cfg.ConnectionStrings.DefaultConnection
		if len(cs1) > 0 {
			fmt.Println(cs1)
		}
		err := godotenv.Load()

		if err != nil {
			log.Println("No .env found")
		}
		dsn := os.Getenv("DB_PG_DSN")

		/*db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		fmt.Println(db)
		if err != nil {
			log.Fatal(err)
		}
		*/
		return dsn
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
