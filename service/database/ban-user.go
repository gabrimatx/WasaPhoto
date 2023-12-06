package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) BanUser(IdUserToBan int, IdUser int) error {
	_, err := db.c.Exec(
		"INSERT INTO Bans(BannerId, BannedId) VALUES (?, ?)", IdUser, IdUserToBan)
	return err
}
