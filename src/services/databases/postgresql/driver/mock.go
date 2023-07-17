package driverpq

import (
	"database/sql"
)

type DBInterface interface {
	Open(driverName string, dataSourceName string) (*sql.DB, error)
	Ping(conn *sql.DB) error
}

type DB struct{}
type MockDB struct {
	DB *sql.DB
	Err error
	ErrPing error
}

func (t MockDB) Open(driverName string, dataSourceName string) (*sql.DB, error) {
	return t.DB, t.Err
}
func (t DB) Open(driverName string, dataSourceName string) (*sql.DB, error) {
	conn, err := sql.Open(driverName, dataSourceName)
	return conn, err
}

func (t MockDB) Ping(conn *sql.DB) (error) {
	if t.ErrPing != nil {
		return t.ErrPing
	}
	return t.Err
}
func (t DB) Ping(conn *sql.DB) (error) {
	err := conn.Ping()
	return err
}