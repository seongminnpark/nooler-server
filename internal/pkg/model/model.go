package model

import (
	"database/sql"
	"errors"
)

type user struct {
	ID    int    `json:"id"`
	UUID  string `json:"uuid"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func (u *user) getUser(db *sql.DB) error {
	return errors.New("getUser not implemented")
}
func (u *user) updateUser(db *sql.DB) error {
	return errors.New("updateUser not implemented")
}
func (u *user) deleteUser(db *sql.DB) error {
	return errors.New("deleteUser not implemented")
}
func (u *user) createUser(db *sql.DB) error {
	return errors.New("createUser not implemented")
}
func getUsers(db *sql.DB, start, count int) ([]user, error) {
	return nil, errors.New("getUsers not implemented")
}
