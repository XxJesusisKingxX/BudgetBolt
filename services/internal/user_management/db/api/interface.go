package api

import (
	"database/sql"

	"services/internal/user_management/db/controller"
	user "services/internal/user_management/db/model"
)

type DBHandler interface {
	RetrieveProfile(db *sql.DB, id string, uid bool) (user.Profile, error)
	CreateProfile(db *sql.DB, user string, password string, randomUID string) error
	CreateToken(db *sql.DB, m user.Token) error
	RetrieveToken(db *sql.DB, profileId int64) (user.Token, error)
}

type DB struct{}
type MockDB struct {
	Profile user.Profile
	Token user.Token
	ProfileErr error
	TokenErr error
}

func (t MockDB) RetrieveProfile(db *sql.DB, id string, uid bool) (user.Profile, error) {
	return t.Profile, t.ProfileErr
}
func (t MockDB) CreateProfile(db *sql.DB, user string, password string, randomUID string) error {
	return t.ProfileErr
}
func (t MockDB) CreateToken(db *sql.DB, m user.Token) error{
	return t.TokenErr
}
func (t MockDB) RetrieveToken(db *sql.DB, profileId int64) (user.Token, error) {
	return t.Token, t.TokenErr
}
func (t DB) RetrieveProfile(db *sql.DB, id string, uid bool) (user.Profile, error) {
	var profile user.Profile;
	var err error;
	if uid {
		profile, err = controller.RetrieveProfile(db, user.Profile{ RandomUID: id })
	} else {
		profile, err = controller.RetrieveProfile(db, user.Profile{ Name: id })
	}
	return profile, err
}
func (t DB) CreateProfile(db *sql.DB, name string, password string, randomUID string) error {
	err := controller.CreateProfile(db, user.Profile{
		Name: name,
		Password: password,
		RandomUID: randomUID,
	})
	return err
}
func (t DB) CreateToken(db *sql.DB, m user.Token) error {
	err := controller.CreateToken(db, m)
	return err
}
func (t DB) RetrieveToken(db *sql.DB, profileId int64) (user.Token, error) {
	token, err := controller.RetrieveToken(db, user.Token{ ProfileID: profileId })
	return token, err
}