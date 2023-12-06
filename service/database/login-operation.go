package database

import (
	"database/sql"
)

func (db *appdbimpl) Login(Username string) (int, error) {
	var Id int
	if err := db.c.QueryRow("SELECT UserId FROM Users WHERE Name = ?", Username).Scan(&Id); err != nil {
		if err == sql.ErrNoRows {
			return -1, err
		}
		return -1, err
	}
	return Id, nil
}
