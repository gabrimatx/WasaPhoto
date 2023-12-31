package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) LikePhoto(IdPhoto uint64, UserLikeId uint64) error {
	_, err := db.c.Exec(
		"INSERT INTO Likes(PhotoId, UserId) VALUES (?, ?)", IdPhoto, UserLikeId)
	return err
}
