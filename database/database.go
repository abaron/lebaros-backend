package database

import (
	"fmt"
	"os"

	"github.com/abaron/lebaros-backend/database/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // configures mysql driver
)

// Initialize initializes the database
func Initialize() (*gorm.DB, error) {
	fmt.Println("\n\n\n", os.Getenv("DB_CONFIG"))
	dbConfig := os.Getenv("DB_CONFIG")
	db, err := gorm.Open("mysql", dbConfig)
	db.LogMode(true) // logs SQL
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	models.Migrate(db)
	return db, err
}
