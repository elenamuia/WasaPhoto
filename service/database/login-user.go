package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) LoginUser(l Login) (UserID int, err error) {
	res, err := db.c.Query(`INSERT OR IGNORE into Users(UserID, username) ((select max(id) from Users)+1,?)`,
		l.UsernameLog)
	if err != nil {
		return 0, err
	}
	res, err = db.c.Query(`select * from Users where username = ?`,
		l.UsernameLog)
	if err != nil {
		return 0, err
	}
	var id int
	for res.Next() {
		var u User
		err = res.Scan(&u.ID, &u.Name)
		if err != nil {
			return 0, err
		}
		id = u.ID
	}
	return id, nil
}
