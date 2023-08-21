package database

func (db *appdbimpl) CheckIfFollow(userid string, idget string) (bool, error) {
	rows, err := db.c.Query(`SELECT Follower, Followed FROM Follower WHERE Follower = ? AND Followed = ? `, userid, idget)
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
