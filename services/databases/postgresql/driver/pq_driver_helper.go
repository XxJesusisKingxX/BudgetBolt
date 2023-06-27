package driverpq

import (
	"golang.org/x/term"
	"syscall"
	"strings"
	"regexp"
	"bufio"
	"fmt"
	"os"
)

func logError(err error, debug bool) {
	if debug {
		fmt.Println(err.Error())
	}
}

func getStdin(msg string, encrypt bool) (string, error) {
	fmt.Print(msg)
	if encrypt {
		encryptInput, _ := term.ReadPassword(int(syscall.Stdin))
		fmt.Println("")
		fmt.Println("") // force a new line after password enter
		return string(encryptInput), nil
	}
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.TrimRight(input, "\n")
	input = strings.TrimSpace(input)
	return input, err
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
