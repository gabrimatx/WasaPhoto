package database

func (db *appdbimpl) DeleteUser(UserId uint64) error {
	_, err := db.c.Exec("DELETE FROM Users WHERE UserId = ?", UserId)
	return err
}
