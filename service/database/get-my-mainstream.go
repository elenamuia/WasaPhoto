package database

func (db *appdbimpl) GetMyMainstream(userid string, page int, perPage int) ([]Photo, error) {
	var ArrayofPhotos []Photo
	offset := (page - 1) * perPage
	rows, err := db.c.Query(`SELECT p.PhotoID, p.User, p.Photo, p.DataPost FROM Photo as p, Follower as f WHERE p.User = f.Followed AND f.Follower = ?  ORDER BY PhotoID DESC LIMIT ? OFFSET ?`, userid, perPage, offset)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var photo Photo
		err = rows.Scan(&photo.ID, &photo.User, &photo.PhotoStructure, &photo.Datapost)
		if err != nil {
			return nil, err
		}

		ArrayofPhotos = append(ArrayofPhotos, photo)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ArrayofPhotos, nil
}
