package main

import (
	"io/ioutil"
	"encoding/json"
	"errors"
	
)

const CONFIG_PATH string = ""

const DB_CONFIG_ENV = CONFIG_PATH + "/db.json"

type dbManager struct {
	dbName   string `json: "db_name"`
	username string `json: "username"`
	password string `json: "password"`
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
 	return !os.IsNotExist(err)
}

func loadDbManager() (dbManager, error) {

	if fileExists(DB_CONFIG_PATH) {
		return , nil

	} else if fileExists(DB_DEFULATS_PATH) {
		return , nil

	} else {
		return nil, errors.New("DB config file does not exist! Make sure you have db.json in /configs.")
	}
}

func main() {
	a := App()

	manager := loadDbManager()

	a.Initialize(manager.username, manager.password, manager.dbName)

	a.Run(":2441")
}