package driverpq

import (
	"budgetbolt/tests"
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

func LogError(err error, debug bool) {
	if debug {
		fmt.Println(err.Error())
	}
}

func getStdin(r io.Reader, msg string, encrypt bool, terminal tests.Terminal) (string, error) {
	fmt.Println(msg)
	if encrypt {
		encryptInput, err := terminal.ReadPassword()
		fmt.Println("") // force a new line after password enter
		if err == nil {
			return string(encryptInput), nil
		}
		return "", err
	}
	readers := bufio.NewReader(r)
	input, err := readers.ReadString('\n')
	if err == nil {
		input = strings.TrimRight(input, "\n")
		input = strings.TrimSpace(input)
		return input, nil
	}
	return "", err
}

func validateUserInput(user string, pass string, db string) bool {
	validUser, _:= regexp.MatchString(`^[^\s%'".,;=(){}[\]$*<>|]+$`, user)
	validPass, _:= regexp.MatchString(`^[^\s]+$`, pass)
	validDB, _:= regexp.MatchString(`^[^\s%'".,;=(){}[\]$*<>|]+$`, db)
	if validUser && validPass && validDB {
		return true
	}
	return false
}
