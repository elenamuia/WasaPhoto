package database

func (db *appdbimpl) GetProfile(user string) (p Profile, err error) {

	p.User = user

	rows1, err1 := db.c.Query(`SELECT p.PhotoID FROM Photo as p WHERE p.User = ?`, user)
	if err1 != nil {

		return p, err1
	}

	for rows1.Next() {

		var photoID int

		err2 := rows1.Scan(&photoID)
		if err2 != nil {

			return p, err
		}

		p.Photos = append(p.Photos, photoID)

	}

	if err5 := rows1.Err(); err5 != nil {
		return p, err
	}
	rows1.Close()

	rows2, err8 := db.c.Query(`SELECT Follower FROM Follower WHERE Followed = ?`, user)
	if err8 != nil {
		return p, err8
	}

	for rows2.Next() {
		var followerid string
		err6 := rows2.Scan(&followerid)
		if err6 != nil {

			return p, err6
		}
		p.Follower = append(p.Follower, followerid)
	}

	if err4 := rows2.Err(); err4 != nil {

		return p, err4
	}
	rows2.Close()

	rows3, err3 := db.c.Query(`SELECT Followed FROM Follower WHERE Follower = ?`, user)
	if err3 != nil {
		return p, err3
	}

	for rows3.Next() {
		var followedid string
		err2 := rows3.Scan(&followedid)
		if err2 != nil {
			return p, err
		}
		p.Followed = append(p.Followed, followedid)

	}

	if err5 := rows3.Err(); err5 != nil {
		return p, err
	}
	rows3.Close()

	return p, nil
}
