package driverpq

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

// Connect to a PostgreSQL database
func connectPQDB(user string, pass string, db string, behaviorDB DBInterface) (*sql.DB, error) {
	loginStr := "user=%v password=%v dbname=%v sslmode=disable"
	connStr := fmt.Sprintf(loginStr, user, pass, db)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = behaviorDB.Ping(conn)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func LogonDB(creds CREDENTIALS, db string, behaviorDB DBInterface, debug bool) (*sql.DB, error) {
	valid := validateUserInput(creds.User, creds.Pass, db)
	if valid {
		conn, err := connectPQDB(creds.User, creds.Pass, db, behaviorDB)
		if err != nil {
			if debug {
				fmt.Println(err.Error())
			}
			return nil, err
		}
		return conn, nil
	}
	return nil, errors.New("Input not valid")
}