package model

import (
	"database/sql"
	"fmt"
)

type Action struct {
	ID     int    `json:"id"`
	UUID   string `json:"uuid"`
	User   string `json:"user"`
	Device string `json:"device"`
	Active bool   `json:"active"`
}

func (action *Action) CreateAction(db *sql.DB) error {

	var active int
	if action.Active {
		active = 1
	} else {
		active = 0
	}

	statement := fmt.Sprintf("INSERT INTO Action(uuid, user, device, active) VALUES('%s', '%s', '%s', b'%d')",
		action.UUID, action.User, action.Device, active)

	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&action.ID)
	if err != nil {
		return err
	}
	return nil
}
