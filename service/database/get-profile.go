package database

func (db *appdbimpl) GetProfile(user string) (p Profile, err error) {

	rows, err1 := db.c.Query("SELECT p.PhotoID, p.User, count(f.Follower) as NumFollowed,  count(f.Followed) as NumFollower FROM Photo as p, Follower as f WHERE User=  ?  GROUP BY p.User ", user)
	if err1 != nil {
		return p, err
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var pr Profile
		err2 := rows.Scan(&pr.User, &pr.Photos, &pr.Followed, &pr.Follower)
		if err2 != nil {
			return p, err
		}

	}
	if err3 := rows.Err(); err3 != nil {
		return p, err
	}

	return p, nil
}
