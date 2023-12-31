package driver

import (
	"services/internal/utils/testing"
	"services/internal/utils/sql/driver"
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

	resValid := driver.ValidateUserInput(userValid, passValid, dbValid)
	resInvalid := driver.ValidateUserInput(userInValid, passInValid, dbInValid)

	// Check if input is valid
	tests.Equals(t, true, resValid)
	// Check if input is invalid
	tests.Equals(t, false, resInvalid)
}

func TestLogError(t *testing.T) {
	err := errors.New("This test is valid")

	isOut := tests.GetStdout(driver.LogError, err, true)
	notOut := tests.GetStdout(driver.LogError, err, false)

	// Check if output is correct
	tests.Equals(t, "This test is valid\n", string(isOut))
	tests.Equals(t, "", string(notOut))

}

func TestConnectPQDB(t *testing.T) {
	db := &sql.DB{}
	err := errors.New("Connection failed")
	errPing := errors.New("Ping failed")
	
	mockDB := driver.MockDB{DB: db}
	mockDBPingFail := driver.MockDB{DB: db, ErrPing: errPing}
	mockDBFail := driver.MockDB{DB: db, Err: err}

	_, caseOne := driver.ConnectPQDB("user", "secret", "test", mockDB)
	_, caseTwo := driver.ConnectPQDB("user", "secret", "test", mockDBFail)
	_, caseThree := driver.ConnectPQDB("user", "secret", "test", mockDBPingFail)

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

	credsFail := driver.CREDENTIALS {User: "$#$@#@", Pass: "%$@#"}
	credsSuccess := driver.CREDENTIALS {User: "user", Pass: "pass"}
	mockDBSuccess := driver.MockDB{DB: db}
	mockDBFail := driver.MockDB{DB: nil, Err: errDBFail}

	_, caseOne := driver.LogonDB(credsFail, "testdb", driver.MockDB{}, false)
	_, caseTwo := driver.LogonDB(credsSuccess, "testdb", mockDBSuccess, false)
	_, caseThree := driver.LogonDB(credsSuccess, "testdb", mockDBFail, false)

	// Check if user doesnt pass validation fails
	tests.EqualsErr(t, err, caseOne)
	// Check if user logon successfully
	tests.Equals(t, nil, caseTwo)
	// Check if user logon fails
	tests.Equals(t, errDBFail, caseThree)
}