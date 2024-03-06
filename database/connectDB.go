// database/connectDB.go

package database

import (
	"github.com/ADEMOLA200/waitlist.git/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDB establishes a connection to the MySQL database
func ConnectDB() *gorm.DB {
    dsn := "root:rootroot@tcp(127.0.0.1:3306)/waitlist?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }
    return db
}

// MigrateDB migrates the database models
func MigrateDB(db *gorm.DB) {
    db.AutoMigrate(&models.Waitlist{}, &models.Campaign{})
}