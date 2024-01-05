package database

import (
	"database/sql"
)

func (db *appdbimpl) GetUser(Username string) (uint64, error) {
	var Id uint64
	err := db.c.QueryRow("SELECT UserId FROM Users WHERE Name = ?", Username).Scan(&Id)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return Id, err

}

func (db *appdbimpl) GetUserName(userId uint64) (string, error) {
	var name string
	err := db.c.QueryRow("SELECT Name FROM Users WHERE UserId = ?", userId).Scan(&name)
	if err == sql.ErrNoRows {
		return "", nil
	}
	return name, err

}
