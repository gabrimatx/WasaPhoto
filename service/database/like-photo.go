package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) LikePhoto(IdPhoto uint64, UserLikeId uint64) error {
	_, err := db.c.Exec(
		"INSERT INTO Likes(PhotoId, UserId) VALUES (?, ?)", IdPhoto, UserLikeId)
	return err
}

func (db *appdbimpl) GetIfLiked(IdPhoto uint64, UserLikeId uint64) (bool, error) {
	var isLiked bool
	row := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM Likes WHERE PhotoId = ? AND UserId = ?)", IdPhoto, UserLikeId)
	if err := row.Scan(&isLiked); err != nil {
		return false, err
	}
	return isLiked, nil
}
