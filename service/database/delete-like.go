package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) DeleteLike(IdPhoto int, UserLikeId int) error {
	_, err := db.c.Exec(
		"DELETE FROM Likes WHERE PhotoId = ? AND UserId = ?", IdPhoto, UserLikeId)
	return err
}
