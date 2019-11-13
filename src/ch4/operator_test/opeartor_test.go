package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestBitClear(t *testing.T) {
	//flag := Readable | Executable
	//t.Log(flag, Readable, Writable, Executable)
	//t.Log(flag&^Readable == Executable)
	//t.Log(flag&^Executable == Readable)
	//t.Log(flag&(^Readable) == Executable)
	//t.Log(flag&(^Executable) == Readable)
	a := 7
	t.Log(a&^Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}

func TestOperator(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	t.Log(a == d)
	//t.Log(a == c)
}
