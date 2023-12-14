package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// Initialize GORM
	initDB()

	// Your other application logic goes here
}

func initDB() {
	// MySQL database configuration
	dsn := "user:password@tcp(127.0.0.1:3306)/database_name?parseTime=true"

	// Open a connection to the database
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Migrate the schema (automatically create tables)
	db.AutoMigrate(&YourModel{}) // Replace YourModel with your actual GORM model
}

// YourModel is an example GORM model
type YourModel struct {
	gorm.Model
	Name string
	// Add other fields as needed
}
