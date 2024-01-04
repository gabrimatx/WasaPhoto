package database

func (db *appdbimpl) GetUserIdFromPhotoId(photoId uint64) (uint64, error) {
	var Id uint64
	err := db.c.QueryRow("SELECT PublisherId FROM Photos WHERE Id = ?", photoId).Scan(&Id)
	return Id, err
}

func (db *appdbimpl) GetUserIdFromCommentId(commentId uint64) (uint64, error) {
	var Id uint64
	err := db.c.QueryRow("SELECT UserId FROM Comments WHERE Id = ?", commentId).Scan(&Id)
	return Id, err
}
