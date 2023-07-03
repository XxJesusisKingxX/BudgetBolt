package tests

import "database/sql"

type MockTerminal struct {
	Password string
	Err      error
}
type MockSql struct {
	DB  *sql.DB
	Err error
	ErrPing error
}

type DB struct {}
type RealTerminal struct{}
type RealSql struct{}

type Terminal interface {
	ReadPassword() (string, error)
}
type Sql interface {
	Open(driverName string, dataSourceName string) (*sql.DB, error)
	Ping(conn *sql.DB) error
}