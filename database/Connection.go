package database

import (
	"log"

	"golang-crud-rest-api/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Init() {
	Connect()
	Migrate()
}

func Connect() {
	Instance, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connect to Database...")
}

func Migrate() {
	_ = Instance.AutoMigrate(&entities.Product{})
	log.Println("Database migration completed ...")
}
