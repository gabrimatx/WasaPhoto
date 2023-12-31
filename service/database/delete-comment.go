package database

func (db *appdbimpl) DeleteComment(commentId uint64) error {
	_, err := db.c.Exec("DELETE FROM Comments WHERE Id = ?", commentId)
	return err
}
