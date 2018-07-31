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

	app := nooler.App{}

	app.Initialize(dbUsername, dbPassword, dbName)

	app.Run(":2441")
}
