package tests

import (
	"database/sql"
	"syscall"

	"golang.org/x/term"
)

func (t MockTerminal) ReadPassword() (string, error) {
	return t.Password, t.Err
}
func (t RealTerminal) ReadPassword() (string, error) {
	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	return string(password), nil
}

func (t MockSql) Open(driverName string, dataSourceName string) (*sql.DB, error) {
	return t.DB, t.Err
}
func (t RealSql) Open(driverName string, dataSourceName string) (*sql.DB, error) {
	conn, err := sql.Open(driverName, dataSourceName)
	return conn, err
}

func (t MockSql) Ping(conn *sql.DB) (error) {
	if t.ErrPing != nil {
		return t.ErrPing
	}
	return t.Err
}
func (t RealSql) Ping(conn *sql.DB) (error) {
	err := conn.Ping()
	return err
}