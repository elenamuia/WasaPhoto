package database

func (db *appdbimpl) GetFollower(userid string) (listFollower []string, err error) {

	rows, err1 := db.c.Query("SELECT Follower FROM Follower WHERE Followed = ? ", userid)
	if err1 != nil {
		return listFollower, err

	}

	for rows.Next() {
		var FollowerId string
		err2 := rows.Scan(&FollowerId)
		if err2 != nil {

			return listFollower, err
		}

		listFollower = append(listFollower, FollowerId)
	}

	return listFollower, nil
}
