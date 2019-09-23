package edb

import "runtime"

// StackTrace gets a trace from the runtime package as a string
func StackTrace() string {
	buf := make([]byte, 2048)
	read := runtime.Stack(buf, false)

	return string(buf[:read])
}
