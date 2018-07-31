package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/seongminnpark/nooler-server/internal/app/nooler"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	a := nooler.App()

	a.Initialize(dbUsername, dbPassword, dbName)

	a.Run(":2441")
}
