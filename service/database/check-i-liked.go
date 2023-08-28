package database

func (db *appdbimpl) CheckILiked(userput string, photoid int) (bool, error) {
	rows, err := db.c.Query(`SELECT UserPutting FROM Like WHERE UserPutting = ? AND PhotoID = ? `, userput, photoid)
	if err != nil {
		return true, err
	}
	defer func() { _ = rows.Close() }()

	if rows.Next() {
		return true, nil
	}

	if err = rows.Err(); err != nil {
		return false, err
	}
	return false, nil
}
