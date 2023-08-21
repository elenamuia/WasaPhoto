package database

func (db *appdbimpl) CheckIfFollow(userid string, idget string) (bool, error) {
	rows, err := db.c.Query(`SELECT Follower, Followed FROM Follower WHERE Follower = ? AND Followed = ? `, userid, idget)
	if err != nil {
		return true, err
	}

	if rows.Next() {
		return true, nil
	}
	rows.Close()
	if err = rows.Err(); err != nil {
		return false, err
	}
	return false, nil
}
