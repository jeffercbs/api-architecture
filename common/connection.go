package common

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=hack4u dbname=apigo port=5432 sslmode=disable TimeZone=America/Bogota"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	return db
}
