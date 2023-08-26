package database

func (db *appdbimpl) CheckILiked(userrec string, photoid int) (bool, error) {
	rows, err := db.c.Query(`SELECT UserReceiving FROM Like WHERE UserReceiving = ? AND PhotoID = ? `, userrec, photoid)
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
