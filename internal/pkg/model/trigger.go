package model

import (
	"database/sql"
	"fmt"
)

type Trigger struct {
	ID     int    `json:"id"`
	UUID   string `json:"uuid"`
	User   string `json:"user"`
	Device string `json:"device"`
}

func (trigger *Trigger) CreateTrigger(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO Trigger(uuid, owner, device) VALUES('%s', '%s', '%s')",
		trigger.UUID, trigger.User, trigger.Device)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&trigger.ID)
	if err != nil {
		return err
	}
	return nil
}
