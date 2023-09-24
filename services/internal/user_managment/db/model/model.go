package model

type Profile struct {
	ID          int  `db:"profile_id"`
	Name      string `db:"profile_name"`
	Password  string `db:"profile_password"`
	RandomUID string `db:"v"`
}
type Token struct {
	ID          int  `db:"token_id"`
	Item      string `db:"item_id"`
	Token     string `db:"access_token"`
	ProfileID    int `db:"profile_id"`
}