package main

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/seongminnpark/nooler-server/internal/app/nooler"
)

var app nooler.App

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	app = nooler.App{}

	app.Initialize(dbUsername, dbPassword, dbName)

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := app.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}
func clearTable() {
	app.DB.Exec("DELETE FROM users")
	app.DB.Exec("ALTER TABLE users AUTO_INCREMENT = 1")
}

const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS users
(
	id INT AUTO_INCREMENT PRIMARY KEY,
	email VARCHAR(50) NOT NULL,
	uuid BINARY(16) NOT NULL,
	token TEXT
)`
