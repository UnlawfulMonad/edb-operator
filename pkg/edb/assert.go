package edb

func Assert(cond bool) {
	if !cond {
		panic("Assertion failed")
	}
}
