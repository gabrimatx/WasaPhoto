package database

func (db *appdbimpl) InsertUser(newUsername string) (int, error) {
	res, err := db.c.Exec(
		`INSERT INTO Users(Name)
         VALUES (?)`, newUsername)
	if err != nil {
		return -1, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(lastInsertID), nil
}
