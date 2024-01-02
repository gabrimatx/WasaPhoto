package database

func (db *appdbimpl) IncrementLikeCount(photoId uint64) error {
	_, err := db.c.Exec("UPDATE Photos SET Likes = Likes + 1 WHERE Id = ?", photoId)
	return err
}

func (db *appdbimpl) DecrementLikeCount(photoId uint64) error {
	_, err := db.c.Exec("UPDATE Photos SET Likes = Likes - 1 WHERE Id = ?", photoId)
	return err
}
