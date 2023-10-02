package model

type Profile struct {
	ID          int64  `json:"profile_id" db:"profile_id"`
	Name        string `json:"profile_name" db:"profile_name"`
	Password    string `json:"profile_password" db:"profile_password"`
	RandomUID   string `json:"v" db:"v"`
}

type Token struct {
	ID          int64    `json:"token_id" db:"token_id"`
	Item        string   `json:"item_id" db:"item_id"`
	Token       string   `json:"access_token" db:"access_token"`
	ProfileID   int64    `json:"profile_id" db:"profile_id"`
}
type Tokens struct {
	Tokens []Token  `json:"tokens"`
}
