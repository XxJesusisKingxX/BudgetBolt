package driverpq

import (
	"budgetbolt/tests"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestIsValidateUserInput(t *testing.T) {
	userValid := "1Qw+!@#&-_`/?"
	userInValid := " %'\".,;=(){}[\\]$*<>|"
	passValid := "1Qw+!@#&-_`/?%'\".,;=(){}[\\]$*<>|]+$"
	passInValid := " "
	dbValid := "1Qw+!@#&-_`/?"
	dbInValid := " %'\".,;=(){}[\\]$*<>|"

	resValid := validateUserInput(userValid, passValid, dbValid)
	resInvalid := validateUserInput(userInValid, passInValid, dbInValid)

	// Check if input is valid
	tests.Equals(t, true, resValid, fmt.Sprintf("Expected %v but got %v", true, resValid))
	// Check if input is invalid
	tests.Equals(t, false, resInvalid, fmt.Sprintf("Expected %v but got %v", false, resInvalid))
}

func TestIsLogError(t *testing.T) {
	err := errors.New("This test is valid")

	isOut := tests.GetStdout(LogError, err, true)
	notOut := tests.GetStdout(LogError, err, false)

	// Check if output is correct
	tests.Equals(t, "This test is valid\n", string(isOut), fmt.Sprintf("Expected %v but got %v", "This test is valid\n", string(isOut)))
	tests.Equals(t, "", string(notOut), fmt.Sprintf("Expected %v but got %v", "", string(notOut)))

}

func TestIsGetStdIn(t *testing.T) {
	err := errors.New("Input error")
	readerValid := strings.NewReader("myusername\n")
	readerInvalid := strings.NewReader("myusername")
	mockTerminal := tests.MockTerminal{Password: "secret123"}
	mockTerminalErr := tests.MockTerminal{Password: "", Err: err}
	realTerminal := tests.RealTerminal{}

	resValid, _ := getStdin(readerValid, "enter username: ", false, realTerminal)
	resInvalid, _  := getStdin(readerInvalid, "enter username: ", false, realTerminal)
	resEncryptValid, _ := getStdin(readerValid, "enter password: ", true, mockTerminal)
	_, resEncryptInvalid := getStdin(readerInvalid, "enter password: ", true, mockTerminalErr)

	// Check if getting user input
	tests.Equals(t, "myusername", resValid, fmt.Sprintf("Expected %v but got %v", "myusername\n", resValid))
	// Check if user input gives error
	tests.Equals(t, "", resInvalid, fmt.Sprintf("Expected %v but got %v", "", resInvalid))
	// Check if user input is encrypt
	tests.Equals(t, "secret123", resEncryptValid, fmt.Sprintf("Expected %v but got %v", "secret123", resEncryptValid))
	// Check if user input is encrypt but gives error
	tests.Equals(t, err, resEncryptInvalid, fmt.Sprintf("Expected %v but got %v", err, resEncryptInvalid))
}

func TestIsConnectPQDB(t *testing.T) {
	db := &sql.DB{}
	err := errors.New("Connection failed")
	errPing := errors.New("Ping failed")
	mockSql := tests.MockSql{DB: db}
	mockSqlPingErr := tests.MockSql{DB: db, ErrPing: errPing}
	mockSqlErr := tests.MockSql{DB: db, Err: err}

	_, connValid := connectPQDB("user", "secret", "test", mockSql)
	_, connInvalid := connectPQDB("user", "secret", "test", mockSqlErr)
	_, connPingInvalid := connectPQDB("user", "secret", "test", mockSqlPingErr)

	// Check if getting connection succeds
	tests.Equals(t, nil, connValid, fmt.Sprintf("Expected %v but got %v", nil, connValid))
	// Check if getting connection fails
	tests.Equals(t, err, connInvalid, fmt.Sprintf("Expected %v but got %v", err, connInvalid))
	// Check if getting connection succeds but ping fails
	tests.Equals(t, errPing, connPingInvalid, fmt.Sprintf("Expected %v but got %v", errPing, connPingInvalid))
}