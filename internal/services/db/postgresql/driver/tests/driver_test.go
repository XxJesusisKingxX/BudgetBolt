package driverpq

import (
	"services/utils/testing"
	"services/db/postgresql/driver"
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

	resValid := driverpq.ValidateUserInput(userValid, passValid, dbValid)
	resInvalid := driverpq.ValidateUserInput(userInValid, passInValid, dbInValid)

	// Check if input is valid
	tests.Equals(t, true, resValid)
	// Check if input is invalid
	tests.Equals(t, false, resInvalid)
}

func TestLogError(t *testing.T) {
	err := errors.New("This test is valid")

	isOut := tests.GetStdout(driverpq.LogError, err, true)
	notOut := tests.GetStdout(driverpq.LogError, err, false)

	// Check if output is correct
	tests.Equals(t, "This test is valid\n", string(isOut))
	tests.Equals(t, "", string(notOut))

}

func TestConnectPQDB(t *testing.T) {
	db := &sql.DB{}
	err := errors.New("Connection failed")
	errPing := errors.New("Ping failed")
	
	mockDB := driverpq.MockDB{DB: db}
	mockDBPingFail := driverpq.MockDB{DB: db, ErrPing: errPing}
	mockDBFail := driverpq.MockDB{DB: db, Err: err}

	_, caseOne := driverpq.ConnectPQDB("user", "secret", "test", mockDB)
	_, caseTwo := driverpq.ConnectPQDB("user", "secret", "test", mockDBFail)
	_, caseThree := driverpq.ConnectPQDB("user", "secret", "test", mockDBPingFail)

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

	credsFail := driverpq.CREDENTIALS {User: "$#$@#@", Pass: "%$@#"}
	credsSuccess := driverpq.CREDENTIALS {User: "user", Pass: "pass"}
	mockDBSuccess := driverpq.MockDB{DB: db}
	mockDBFail := driverpq.MockDB{DB: nil, Err: errDBFail}

	_, caseOne := driverpq.LogonDB(credsFail, "testdb", driverpq.MockDB{}, false)
	_, caseTwo := driverpq.LogonDB(credsSuccess, "testdb", mockDBSuccess, false)
	_, caseThree := driverpq.LogonDB(credsSuccess, "testdb", mockDBFail, false)

	// Check if user doesnt pass validation fails
	tests.EqualsErr(t, err, caseOne)
	// Check if user logon successfully
	tests.Equals(t, nil, caseTwo)
	// Check if user logon fails
	tests.Equals(t, errDBFail, caseThree)
}