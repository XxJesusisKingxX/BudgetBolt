package driverpq

import (
	"budgetbolt/tests"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// Connect to a PostgreSQL database
func connectPQDB(user string, pass string, db string) (*sql.DB, error) {
	loginStr := "user=%v password=%v dbname=%v sslmode=disable"
	connStr := fmt.Sprintf(loginStr, user, pass, db)
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
	db, dbErr := getStdin(os.Stdin, "Enter database name: ", false, tests.RealTerminal{})
	user, userErr := getStdin(os.Stdin, "Enter username: ", false, tests.RealTerminal{})
	pass, passErr := getStdin(os.Stdin, "Enter password: ", true, tests.RealTerminal{})
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
		LogError(dbErr, debug)
	}
	if userErr != nil {
		LogError(userErr, debug)
	}
	if passErr != nil {
		LogError(passErr, debug)
	}
	return nil, err
}