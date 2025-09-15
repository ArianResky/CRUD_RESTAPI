package main

import (
	"log"
	"os"

	"crud_restapi/repository"
	"crud_restapi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	routes.RegisterRoutes(r, db)

	addr := getEnv("ADDR", ":8080")
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
