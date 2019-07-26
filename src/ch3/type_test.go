package type_test

import "testing"

func TestImplicit(t *testing.T) {
	// 不支持隐式类型转换
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

type MyInt int64

func TestPoint(t *testing.T) {
	a := 1
	// 指针不支持运算
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	// string 值类型 初始化为空值串 而不是nil
	var s string
	t.Log("#" + s + "#")
	t.Log(len(s))
}
