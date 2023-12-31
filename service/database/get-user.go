package database

import (
	"database/sql"
)

func (db *appdbimpl) GetUser(Username string) (uint64, error) {
	var Id uint64
	if err := db.c.QueryRow("SELECT UserId FROM Users WHERE Name = ?", Username).Scan(&Id); err != nil {
		if err == sql.ErrNoRows {
			return 0, err
		}
		return 0, err
	}
	return Id, nil
}
