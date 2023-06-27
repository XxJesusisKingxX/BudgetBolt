package tests

import (
	"testing"
)

func Equals(t *testing.T, a any, b any, msg string) {
	if a != b {
		t.Helper()
		t.Errorf(msg)
	}
}