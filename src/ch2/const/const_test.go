package _const

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConst(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConst1(t *testing.T) {
	a := 7 // 0111
	t.Log(a&Readable, a&Writable, a&Executable)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
