package edb

import (
	"regexp"
	"runtime"

	"github.com/sethvargo/go-password/password"
)

var (
	userValidateRegexp = regexp.MustCompile(`^[a-z][a-z0-9]*$`)
)

// Checks if the passed string is a valid identifier.
func isValidIdentifier(name string) bool {
	return userValidateRegexp.MatchString(name)
}

// StackTrace gets a trace from the runtime package as a string
func StackTrace() string {
	buf := make([]byte, 2048)
	read := runtime.Stack(buf, false)

	return string(buf[:read])
}

// GenPassword generates a 64 character password.
// May panic on error calling `password.Generate`.
func GenPassword() string {
	result, err := password.Generate(64, 10, 10, false, true)
	if err != nil {
		panic(err)
	}

	return result
}
