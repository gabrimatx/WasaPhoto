package database

func (db *appdbimpl) DeleteFollow(IdUserToNotFollow uint64, IdFollowingUser uint64) error {
	_, err := db.c.Exec(
		"DELETE FROM Follows WHERE FollowerId = ? AND FollowedId = ?", IdFollowingUser, IdUserToNotFollow)
	return err
}
