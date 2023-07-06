package tests

import (
	"io"
	"os"
	"reflect"
)

func GetStdout(f interface{}, args ...interface{}) string {
	// swap the stdout
	tmp := os.Stdout
	// capture the stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	// Call the function using reflection
	fn := reflect.ValueOf(f)
	fnArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		fnArgs[i] = reflect.ValueOf(arg)
	}
	fn.Call(fnArgs)
	w.Close()
	out, _ := io.ReadAll(r) // read stdout
	// swap back stdout
	os.Stdout = tmp
	return string(out)
}


