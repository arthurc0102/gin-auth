package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Load sqlite driver
)

// Connection to database
var Connection *gorm.DB

// Connect to database
func Connect() {
	dbConfig := os.Getenv("DB_CONFIG")
	if dbConfig == "" {
		log.Fatalln("No database config set")
	}

	db, err := gorm.Open("sqlite3", dbConfig)
	if err != nil {
		log.Fatalln(err)
	}

	db.SingularTable(true)
	Connection = db
}

// Close db connections
func Close() {
	Connection.Close()
}
