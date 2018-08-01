package model

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID    int    `json:"id"`
	UUID  string `json:"uuid"`
	Email string `json:"email"`
}

func (user *User) GetUser(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT email FROM users WHERE uuid=%s", user.UUID)
	return db.QueryRow(statement).Scan(&user.Email)
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

func GetUsers(db *sql.DB, start, count int) ([]User, error) {
	statement := fmt.Sprintf("SELECT id, email FROM users LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []User{}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
