package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) LikePhoto(IdPhoto int, UserLikeId int) error {
	_, err := db.c.Exec(
		"INSERT INTO Likes(PhotoId, UserId) VALUES (?, ?)", IdPhoto, UserLikeId)
	return err
}
