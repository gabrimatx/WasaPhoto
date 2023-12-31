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
