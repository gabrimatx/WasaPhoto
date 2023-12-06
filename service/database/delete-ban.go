package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) DeleteBan(IdUserToUnban int, IdUser int) error {
	_, err := db.c.Exec(
		"DELETE FROM Follows WHERE BannerId = ? AND BannedId = ?", IdUser, IdUserToUnban)
	return err
}
