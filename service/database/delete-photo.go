package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) DeletePhoto(id int) error {
	_, err := db.c.Exec("DELETE FROM Photos WHERE Id = ?", id)
	return err
}
