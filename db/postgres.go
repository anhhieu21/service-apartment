package db

import (
	"apartment/internal/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DatabasaConnection() {
	host := "172.17.0.2"
	port := "5432"
	dbName := "hieudev"
	dbUser := "hieudev"
	password := "hieudev21"
	dns := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		host,
		port,
		dbName,
		dbUser,
		password)

	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	DB.AutoMigrate(models.Apartment{})
	DB.AutoMigrate(models.Customer{})
	if err != nil {
		log.Fatal("Error connectiong to the database...", err)
	}
	fmt.Println("Databse connection successful")
}
