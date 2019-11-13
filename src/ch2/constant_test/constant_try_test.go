package constant_test

import "testing"

const(
	Monday = iota + 1
	Tuesday
	Wednesday
)

const(
	Readbale = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstantTry1(t *testing.T) {
	a := 7
	t.Log(Readbale, Writable, Executable)
	t.Log(a&Readbale==Readbale, a&Writable==Writable, a&Executable==Executable)
}
