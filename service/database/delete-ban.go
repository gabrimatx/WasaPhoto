package database

func (db *appdbimpl) DeleteBan(IdUserToUnban uint64, IdUser uint64) error {
	_, err := db.c.Exec(
		"DELETE FROM Bans WHERE BannerId = ? AND BannedId = ?", IdUser, IdUserToUnban)
	return err
}
