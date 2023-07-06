package tests

import (
	"errors"
	"testing"
)

func Equals(t *testing.T, a any, b any, msg string) {
	if a != b {
		t.Helper()
		t.Errorf(msg)
	}
}

func EqualsErr(t *testing.T, a error, b error, msg string) {
	if errors.Is(a, b){
		t.Helper()
		t.Errorf(msg)
	}
}