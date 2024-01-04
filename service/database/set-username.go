package database

func (db *appdbimpl) SetUsername(UserId uint64, new_username string) error {
	_, err := db.c.Exec(
		`UPDATE Users 
	     SET Name = ?
		 WHERE UserId = ?`, new_username, UserId)
	return err
}
