package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
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

	dbName := os.Getenv("TEST_DB_NAME")
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
	if _, err := app.DB.Exec(userTableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	app.DB.Exec("DELETE FROM users")
	app.DB.Exec("ALTER TABLE users AUTO_INCREMENT = 1")
}

const userTableCreationQuery = `
CREATE TABLE IF NOT EXISTS User
(
	id INT AUTO_INCREMENT PRIMARY KEY,
	email VARCHAR(50) NOT NULL,
	uuid VARCHAR(36) NOT NULL,
	password_hash VARCHAR(64) NOT NULL  
)`

const deviceTableCreationQuery = `
CREATE TABLE IF NOT EXISTS Device
(
	id INT AUTO_INCREMENT PRIMARY KEY,
	owner VARCHAR(36) NOT NULL,
	uuid VARCHAR(36) NOT NULL
)`

const actionTableCreationQuery = `
CREATE TABLE IF NOT EXISTS Action
(
	id INT AUTO_INCREMENT PRIMARY KEY,
	uuid VARCHAR(36) NOT NULL,
	user VARCHAR(36) NOT NULL,
	device VARCHAR(36) NOT NULL,
	active BIT NOT NULL
)`

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/users", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
func TestGetNonExistentUser(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/user", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusUnauthorized, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["error"] != "Invalid token" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Invalid token'. Got '%s'", m["error"])
	}
}

func TestCreateUser(t *testing.T) {
	clearTable()

	payload := []byte(`{"email":"test@test.com", "password":"password"}`)

	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	// var m map[string]interface{}
	// json.Unmarshal(response.Body.Bytes(), &m)

	// if m["token"] != "test@test.com" {
	// 	t.Errorf("Expected user email to be 'test@test.com'. Got '%v'", m["email"])
	// }
	// if m["uuid"] == nil {
	// 	t.Errorf("Expected user uuid to be non empty.")
	// }
	// // the id is compared to 1.0 because JSON unmarshaling converts numbers to
	// // floats, when the target is a map[string]interface{}
	// if m["id"] != 1.0 {
	// 	t.Errorf("Expected user ID to be '1'. Got '%v'", m["id"])
	// }
}

func TestGetUser(t *testing.T) {
	clearTable()

	if err := addUsers(1); err != nil {
		t.Errorf("Error while adding users.")
	}

	req, _ := http.NewRequest("GET", "/user", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func addUsers(count int) error {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		stringI := strconv.Itoa(i)
		email := "user" + stringI + "@test.com"
		uuid := "Uuid" + stringI
		statement := fmt.Sprintf("INSERT INTO users(email, uuid) VALUES('%s', '%s')", email, uuid)

		if _, err := app.DB.Exec(statement); err != nil {
			return err
		}
	}

	return nil
}

func TestUpdateUser(t *testing.T) {
	clearTable()

	if err := addUsers(1); err != nil {
		t.Errorf("Error while adding users.")
	}

	req, _ := http.NewRequest("GET", "/user/1", nil)
	response := executeRequest(req)

	var originalUser map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalUser)

	payload := []byte(`{"email":"updated@test.com"}`)
	req, _ = http.NewRequest("PUT", "/user/1", bytes.NewBuffer(payload))
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != originalUser["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalUser["id"], m["id"])
	}
	if m["uuid"] != originalUser["uuid"] {
		t.Errorf("Expected the uuid to remain the same (%v). Got %v", originalUser["uuid"], m["uuid"])
	}
	if m["email"] == originalUser["email"] {
		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalUser["email"], m["email"], m["email"])
	}
}

func TestDeleteUser(t *testing.T) {
	clearTable()

	if err := addUsers(1); err != nil {
		t.Errorf("Error while adding users.")
	}

	req, _ := http.NewRequest("GET", "/user", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/user", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/user", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}
