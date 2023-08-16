package database

func (db *appdbimpl) GetBannedList(userid string) (listban []string, err error) {

	rows, err1 := db.c.Query("SELECT Banning FROM Banned WHERE Banned = ? ", userid)
	if err1 != nil {
		return listban, err
	}

	for rows.Next() {
		var ban string
		err2 := rows.Scan(&ban)
		if err2 != nil {

			return listban, err
		}

		listban = append(listban, ban)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return listban, nil
}
