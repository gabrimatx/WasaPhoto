package database

import (
	components "github.com/gabrimatx/WasaPhoto/service"
)

func (db *appdbimpl) DeleteUser(UserId uint64) error {
	res, err := db.c.Exec("DELETE FROM Users WHERE UserId = ?", UserId)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the user didn't exist
		return components.ErrUserNotExists
	}
	return nil
}
