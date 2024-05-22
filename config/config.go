package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() *gorm.DB {

	connStr := "postgresql://postgres:0745@localhost/applicationTracker?sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db

}
