package database

import (
	"log"

	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func GetDBInstance() *gorm.DB {
	dbURL := os.Getenv("DBURL")

	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
