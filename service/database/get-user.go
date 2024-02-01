package database

import (
	"database/sql"

	components "github.com/gabrimatx/WasaPhoto/service"
)

func (db *appdbimpl) GetUser(Username string) (uint64, error) {
	var Id uint64
	err := db.c.QueryRow("SELECT UserId FROM Users WHERE Name = ?", Username).Scan(&Id)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return Id, err

}

func (db *appdbimpl) GetUserSearch(Username string) (components.UserSearchList, error) {
	var ToReturn components.UserSearchList
	rows, err := db.c.Query("SELECT UserId, Name FROM Users WHERE Name LIKE ?", "%"+Username+"%")
	if err != nil {
		return ToReturn, err
	}
	defer rows.Close()

	for rows.Next() {
		var TempUser components.UserSearchElement
		if err := rows.Scan(&TempUser.Id, &TempUser.Username); err != nil {
			return ToReturn, err
		}
		ToReturn.UList = append(ToReturn.UList, TempUser)
	}

	if err := rows.Err(); err != nil {
		return ToReturn, err
	}
	return ToReturn, err
}

func (db *appdbimpl) GetUserName(userId uint64) (string, error) {
	var name string
	err := db.c.QueryRow("SELECT Name FROM Users WHERE UserId = ?", userId).Scan(&name)
	if err == sql.ErrNoRows {
		return "", nil
	}
	return name, err

}
