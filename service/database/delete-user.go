package database

func (db *appdbimpl) DeleteUser(UserId int) error {
	_, err := db.c.Exec("DELETE FROM Users WHERE Id = ?", UserId)
	return err
}
