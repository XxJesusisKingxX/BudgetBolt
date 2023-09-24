package utils

import (
	"database/sql"

	ctrl "services/internal/user_managment/db/controller"
	"services/internal/user_managment/db/model"
)

func ProfileToDB(db *sql.DB, username string, password string){
	profile := model.Profile {
		Name: username,
		Password: password,
	}
	ctrl.CreateProfile(db, profile)
}

func TokenToDB(db *sql.DB, itemId string, accessToken string){
	token := model.Token {
		Item: itemId,
		Token: accessToken,
	}
	ctrl.CreateToken(db, token)
}