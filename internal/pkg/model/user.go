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

type SignupForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) GetUser(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT email, password_hash FROM User WHERE uuid='%s'", user.UUID)
	return db.QueryRow(statement).Scan(&user.Email, &user.PasswordHash)
}

func (user *User) GetUserByEmail(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT uuid, password_hash FROM User WHERE email='%s'", user.Email)
	return db.QueryRow(statement).Scan(&user.UUID, &user.PasswordHash)
}

func (user *User) UpdateUser(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE User SET email='%s' WHERE uuid='%s'", user.Email, user.UUID)
	_, err := db.Exec(statement)
	return err
}

func (user *User) DeleteUser(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM User WHERE uuid='%s'", user.UUID)
	_, err := db.Exec(statement)
	return err
}

func (user *User) CreateUser(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO User(email, uuid, password_hash) VALUES('%s', '%s', '%s')",
		user.Email, user.UUID, user.PasswordHash)
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
