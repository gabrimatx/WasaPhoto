package database

func (db *appdbimpl) DeleteComment(commentId int) error {
	_, err := db.c.Exec("DELETE FROM Comments WHERE Id = ?", commentId)
	return err
}
