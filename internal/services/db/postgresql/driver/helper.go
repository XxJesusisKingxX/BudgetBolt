package driverpq

import (
	"regexp"
	"fmt"
)

func LogError(err error, debug bool) {
	if debug {
		fmt.Println(err.Error())
	}
}

func ValidateUserInput(user string, pass string, db string) bool {
	validUser, _:= regexp.MatchString(`^[^\s%'".,;=(){}[\]$*<>|]+$`, user)
	validPass, _:= regexp.MatchString(`^[^\s]+$`, pass)
	validDB, _:= regexp.MatchString(`^[^\s%'".,;=(){}[\]$*<>|]+$`, db)
	if validUser && validPass && validDB {
		return true
	}
	return false
}
