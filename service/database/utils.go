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

func (db *appdbimpl) GetBoolFollow(userId uint64, myId uint64) (bool, error) {
	row := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM Follows WHERE FollowerId = ? AND FollowedId = ?)", myId, userId)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}

func (db *appdbimpl) GetBoolBanned(userId uint64, myId uint64) (bool, error) {
	row := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM Bans WHERE BannerId = ? AND BannedId = ?)", myId, userId)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}
