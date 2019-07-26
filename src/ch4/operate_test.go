package ch4

import "testing"

// go 没有前置 ++, -- 有后置+=, --
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 5}
	b := [...]int{1, 2, 5, 3}
	c := [...]int{1, 2, 3, 4}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == b)
	t.Log(a == c)
	// true go 数组维度相同且长度相等时 可以==比较(不是比较引用。。) 每个元素相同才相等
	t.Log(c == d)
}

const (
	Readable = 1 << iota
	Writable
	Executable
)

// &^ 按位清零
func TestBitClear(t *testing.T) {
	// 右边操作数是1时, 左边操作数清零。 右边操作数是0, 左操作数不变

	a := 7 // 0111
	a = a &^ Readable
	t.Log(a&Readable, a&Writable, a&Executable)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
