package database

func (db *appdbimpl) GetFollowed(userid string) (listFollowed []string, err error) {

	rows, err1 := db.c.Query("SELECT Followed FROM Follower WHERE Follower = ? ", userid)
	if err1 != nil {
		return listFollowed, err

	}

	for rows.Next() {
		var FollowedId string
		err2 := rows.Scan(&FollowedId)
		if err2 != nil {

			return listFollowed, err
		}

		listFollowed = append(listFollowed, FollowedId)
	}

	return listFollowed, nil
}
