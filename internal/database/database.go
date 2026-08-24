package database

import (
	"fmt"
	"os"

	"github.com/insignificantGuy/Slotly/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// parseTime=True is required for MySQL DATETIME columns to scan into time.Time.
const defaultDSN = "root@tcp(127.0.0.1:3306)/slotly?charset=utf8mb4&parseTime=True&loc=Local"

func InitMySQL() {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		dsn = defaultDSN
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}
}
