package tests

import (
	"errors"
	"testing"
)

func Equals(t *testing.T, a any, b any) {
	if a != b {
		t.Helper()
		t.Errorf("Expected %v but got %v", a, b)
	}
}

func EqualsErr(t *testing.T, a error, b error) {
	if errors.Is(a, b){
		t.Helper()
		t.Errorf("Expected %v but got %v", a, b)
	}
}