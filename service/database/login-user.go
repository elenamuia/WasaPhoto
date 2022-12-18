package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) LoginUser(UserName string) (UserID int, err error) {
	res, err := db.c.Exec(`INSERT OR IGNORE into USERS(UserID, username) ((select max(UserID)fromUSERS))+ 1,?)`, UserName)
	if err != nil {
		return 0, err
	}

	res, err = db.c.Query(`SELECT * FROM Users where username =?`, UserName)
	if err != nil {
		return UserID, err
	}

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
