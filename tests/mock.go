package tests

import (
	"syscall"

	"golang.org/x/term"
)

func (t MockTerminal) ReadPassword() (string, error) {
	return t.Password, t.Err
}

func (t RealTerminal) ReadPassword() (string, error) {
	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	return string(password), nil
}