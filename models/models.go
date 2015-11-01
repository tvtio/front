package models

import (
	"database/sql"
	// load MySQL driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/tvtio/front/logger"
)

// GetUser a user
func GetUser(id string) (user *User, err error) {
	db, err := sql.Open("mysql", "root:C4lldC@tcp(tvt.io:3306)/tvtio?charset=utf8")
	defer db.Close()
	if err != nil {
		logger.Fatal(err.Error())
	}

	if id != "<nil>" {
		rows, err := db.Query("SELECT * FROM users where id = ?", id)
		if err != nil {
			logger.Fatal(err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var (
				id   string
				name string
			)
			err = rows.Scan(&id, &name)
			if err != nil {
				logger.Fatal(err.Error())
			}
			var user User
			user.ID = id
			user.Name = name
			return &user, nil
		}
		err = rows.Err()
		if err != nil {
			logger.Fatal(err.Error())
		}

		return user, nil
	}
	return nil, err
}
