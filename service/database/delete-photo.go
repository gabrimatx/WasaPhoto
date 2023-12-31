package database

func (db *appdbimpl) DeletePhoto(id uint64) error {
	_, err := db.c.Exec("DELETE FROM Photos WHERE Id = ?", id)
	return err
}
