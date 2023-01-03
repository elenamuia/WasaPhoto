package database

func (db *appdbimpl) GetPhoto() (photo Photo, err error) {
	var PhotoID int
	err = db.c.QueryRow("SELECT * FROM Photo WHERE PhotoID=  ? ").Scan(&PhotoID)
	return photo, err
}
