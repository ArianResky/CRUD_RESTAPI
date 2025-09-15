package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

  "crud_restapi/models"
)


func ConnectDB() (*gorm.DB, error) {
	host := getEnv("DB_HOST", "127.0.0.1")
	port := getEnv("DB_PORT", "3306")
	user := getEnv("DB_USER", "appuser")
	pass := getEnv("DB_PASS", "supersecret")
	name := getEnv("DB_NAME", "myappdb")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&charset=utf8mb4",
		user, pass, host, port, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Println("failed to connect to MariaDB:", err)
		return nil, err
	}

	if err := db.AutoMigrate(&models.Book{}); err != nil {
		log.Println("failed to migrate:", err)
		return nil, err
	}

	return db, nil
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
