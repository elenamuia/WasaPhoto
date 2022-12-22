package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) LoginUser(l Login) (UserID int, err error) {
	_, err = db.c.Exec(`INSERT OR IGNORE into Users(UserID, username) ((select max(id) from Users)+1,?)`,
		l.UsernameLog)
	if err != nil {
		return 0, err
	}
	rows, err := db.c.Query(`select * from Users where username = ?`, l.UsernameLog)
	if err != nil {
		return 0, err
	}
	var id int
	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Name)
		if err != nil {
			return 0, err
		}
		id = u.ID
	}
	if err = rows.Err(); err != nil {
		return 0, err
	}
	return id, nil
}
