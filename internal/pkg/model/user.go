package model

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID           int    `json:"id"`
	UUID         string `json:"uuid"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

func (user *User) GetUser(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT email, password_hash FROM users WHERE uuid=%s", user.UUID)
	return db.QueryRow(statement).Scan(&user.Email, &user.PasswordHash)
}

func (user *User) GetUserByEmail(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT uuid, password_hash FROM users WHERE email=%s", user.Email)
	return db.QueryRow(statement).Scan(&user.UUID, &user.PasswordHash)
}

func (user *User) UpdateUser(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE users SET email='%s' WHERE uuid=%s", user.Email, user.UUID)
	_, err := db.Exec(statement)
	return err
}

func (user *User) DeleteUser(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM users WHERE uuid=%s", user.UUID)
	_, err := db.Exec(statement)
	return err
}

func (user *User) CreateUser(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO users(email, uuid) VALUES('%s', '%s')",
		user.Email, user.UUID)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}
