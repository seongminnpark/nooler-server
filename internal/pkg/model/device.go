package model

import (
	"database/sql"
	"fmt"
)

type Device struct {
	ID       int    `json:"id"`
	UUID     string `json:"uuid"`
	Owner    string `json:"owner"`
	Location string `json:"location"`
}

type AddDeviceForm struct {
	Owner    string `json:"owner"`
	Location string `json:"location"`
}

func (device *Device) GetDevice(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT owner, location FROM Device WHERE uuid='%s'", device.UUID)
	return db.QueryRow(statement).Scan(&device.Owner, &device.Location)
}

func (device *Device) UpdateDevice(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE Device SET owner='%s', location='%s' WHERE uuid='%s'",
		device.Owner, device.Location, device.UUID)
	_, err := db.Exec(statement)
	return err
}

func (device *Device) DeleteDevice(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM Device WHERE uuid='%s'", device.UUID)
	_, err := db.Exec(statement)
	return err
}

func (device *Device) CreateDevice(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO Device(uuid, owner, location) VALUES('%s', '%s', '%s')",
		device.UUID, device.Owner, device.Location)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&device.ID)
	if err != nil {
		return err
	}
	return nil
}
