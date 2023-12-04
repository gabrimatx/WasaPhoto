package database

func (db *appdbimpl) DeletePhoto(id int) error {
	_, err := db.c.Exec("DELETE FROM Photos WHERE Id = ?", id)
	return err
}
