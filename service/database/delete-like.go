package database

func (db *appdbimpl) DeleteLike(IdPhoto uint64, UserLikeId uint64) error {
	_, err := db.c.Exec("DELETE FROM Likes WHERE PhotoId = ? AND UserId = ?", IdPhoto, UserLikeId)
	return err
}
