package tests

type MockTerminal struct {
	Password string
	Err      error
}

type RealTerminal struct{}

type Terminal interface {
	ReadPassword() (string, error)
}