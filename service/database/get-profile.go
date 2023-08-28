package database

import (
	"database/sql"
	"sort"
)

func (db *appdbimpl) GetProfile(user string) (p Profile, err error) {

	err6 := db.c.QueryRow(`SELECT Name from Users WHERE Name =?`, user).Scan(&p.User)
	if err6 != nil {
		if err6 == sql.ErrNoRows {
			return p, err
		}
	}

	rows1, err1 := db.c.Query(`SELECT p.PhotoID FROM Photo as p WHERE p.User = ?`, user)
	if err1 != nil {

		return p, err1
	}
	defer func() { _ = rows1.Close() }()

	for rows1.Next() {

		var photoID int

		err2 := rows1.Scan(&photoID)
		if err2 != nil {

			return p, err
		}

		p.Photos = append(p.Photos, photoID)

	}
	sort.Sort(sort.Reverse(sort.IntSlice(p.Photos)))

	if err5 := rows1.Err(); err5 != nil {
		return p, err
	}

	rows2, err8 := db.c.Query(`SELECT Follower FROM Follower WHERE Followed = ?`, user)
	if err8 != nil {
		return p, err8
	}
	defer func() { _ = rows2.Close() }()

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

	rows3, err3 := db.c.Query(`SELECT Followed FROM Follower WHERE Follower = ?`, user)
	if err3 != nil {
		return p, err3
	}
	defer func() { _ = rows3.Close() }()

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

	return p, nil
}
