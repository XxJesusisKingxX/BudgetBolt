package driverpq

import (
	"budgetbolt/tests"
	"testing"
	"fmt"
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