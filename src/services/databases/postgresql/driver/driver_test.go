package driverpq

import (
	"budgetbolt/src/services/tests"
	"database/sql"
	"testing"
	"errors"
)

func TestValidateUserInput(t *testing.T) {
	userValid := "1Qw+!@#&-_`/?"
	userInValid := " %'\".,;=(){}[\\]$*<>|"
	passValid := "1Qw+!@#&-_`/?%'\".,;=(){}[\\]$*<>|]+$"
	passInValid := " "
	dbValid := "1Qw+!@#&-_`/?"
	dbInValid := " %'\".,;=(){}[\\]$*<>|"

	resValid := validateUserInput(userValid, passValid, dbValid)
	resInvalid := validateUserInput(userInValid, passInValid, dbInValid)

	// Check if input is valid
	tests.Equals(t, true, resValid)
	// Check if input is invalid
	tests.Equals(t, false, resInvalid)
}

func TestLogError(t *testing.T) {
	err := errors.New("This test is valid")

	isOut := tests.GetStdout(LogError, err, true)
	notOut := tests.GetStdout(LogError, err, false)

	// Check if output is correct
	tests.Equals(t, "This test is valid\n", string(isOut))
	tests.Equals(t, "", string(notOut))

}

func TestConnectPQDB(t *testing.T) {
	db := &sql.DB{}
	err := errors.New("Connection failed")
	errPing := errors.New("Ping failed")
	
	mockDB := MockDB{DB: db}
	mockDBPingFail := MockDB{DB: db, ErrPing: errPing}
	mockDBFail := MockDB{DB: db, Err: err}

	_, caseOne := connectPQDB("user", "secret", "test", mockDB)
	_, caseTwo := connectPQDB("user", "secret", "test", mockDBFail)
	_, caseThree := connectPQDB("user", "secret", "test", mockDBPingFail)

	// Check if getting connection succeds
	tests.Equals(t, nil, caseOne)
	// Check if getting connection fails
	tests.Equals(t, err, caseTwo)
	// Check if getting connection succeds but ping fails
	tests.Equals(t, errPing, caseThree)
}

func TestLogonDB(t *testing.T) {
	db := &sql.DB{}
	err := errors.New("Input not valid")
	errDBFail := errors.New("Connection failed")

	credsFail := CREDENTIALS {User: "$#$@#@", Pass: "%$@#"}
	credsSuccess := CREDENTIALS {User: "user", Pass: "pass"}
	mockDBSuccess := MockDB{DB: db}
	mockDBFail := MockDB{DB: nil, Err: errDBFail}

	_, caseOne := LogonDB(credsFail, "testdb", MockDB{}, false)
	_, caseTwo := LogonDB(credsSuccess, "testdb", mockDBSuccess, false)
	_, caseThree := LogonDB(credsSuccess, "testdb", mockDBFail, false)

	// Check if user doesnt pass validation fails
	tests.EqualsErr(t, err, caseOne)
	// Check if user logon successfully
	tests.Equals(t, nil, caseTwo)
	// Check if user logon fails
	tests.Equals(t, errDBFail, caseThree)
}