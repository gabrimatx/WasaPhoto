package database

func (db *appdbimpl) GetNumberOfPhotos(UserId uint64) (int, error) {
	var num int
	err := db.c.QueryRow(
		`
		SELECT COUNT(*) 
		FROM Photos 
		WHERE PublisherId = ?
		`, UserId).Scan(&num)

	if err != nil {
		return -1, err
	}

	return num, err
}

func (db *appdbimpl) PhotoCascadeDeletion(photoId uint64) error {
	_, err := db.c.Exec("DELETE FROM Likes WHERE PhotoId = ?", photoId)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("DELETE FROM Comments WHERE PhotoId = ?", photoId)
	if err != nil {
		return err
	}

	return err
}
