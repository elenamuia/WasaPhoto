package database

func (db *appdbimpl) GetPagePhoto(userid string, page int, perPage int) (photos []Photo, err error) {
	offset := (page - 1) * perPage

	rows, err1 := db.c.Query("SELECT PhotoID, User,Photo FROM Photo WHERE User=? LIMIT ? OFFSET ? ", userid, perPage, offset)

	if err1 != nil {
		return nil, err1
	}
	defer rows.Close()

	for rows.Next() {
		var p Photo
		if err2 := rows.Scan(&p.ID, &p.User, &p.PhotoStructure); err2 != nil {
			return nil, err2
		}

		photos = append(photos, p)

	}

	if err5 := rows.Err(); err5 != nil {
		return photos, err5
	}

	return photos, nil
}
