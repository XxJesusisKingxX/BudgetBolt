package driverpq

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
)

// Connect to a PostgreSQL database
func connectPQDB(user string, pass string, db string) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", user, pass, db)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func LogonDB(debug bool) (*sql.DB, error) {
	db, dbErr := getStdin("Enter database name: ", false)
	user, userErr := getStdin("Enter username: ", false)
	pass, passErr := getStdin("Enter password: ", true)
	if dbErr == nil && userErr == nil && passErr == nil {
		valid := validateUserInput(user, pass, db)
		if valid {
			conn, err := connectPQDB(user, pass, db)
			if err != nil {
				if debug {
					fmt.Println(err.Error())
				}
				return nil, err
			}
			return conn, nil
		}
	}
	var err error
	if dbErr != nil {
		logError(dbErr, debug)
		return nil, dbErr
	} else if userErr != nil {
		logError(userErr, debug)
		return nil, userErr
	} else if passErr != nil {
		logError(passErr, debug)
		return nil, passErr
	}
	return nil, err
}