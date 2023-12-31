package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) BanUser(IdUserToBan uint64, IdUser uint64) error {
	_, err := db.c.Exec(
		"INSERT INTO Bans(BannerId, BannedId) VALUES (?, ?)", IdUser, IdUserToBan)
	return err
}
