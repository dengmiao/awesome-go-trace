package ch9

import "testing"

// string是数据类型 不是引用或指针类型
// string是只读的byte slice, len函数可以获取它所包含的byte数
// string的byte数组可以存放任何数据
func TestString(t *testing.T) {
	var s string
	// 初始化零值 ""
	t.Log("->" + s + "<-")
	s = "hello"
	t.Log(len(s))
	// string是不可变的byte slice 尝试修改 编译错误
	//s[1] = '3'
	// 可存放任意的二进制数据
	s = "\xE4\xB8\xA5"
	// len 求出的是byte数
	t.Log(s, len(s))
	s = "\xE4\xBA\xBB\xFF"
	t.Log(s, len(s))

	s = "中"
	// byte数
	t.Log(len(s))

	// rune取出字符串的unicode
	c := []rune(s)
	t.Log(len(c))
	// Unicode是一种字符集 UTF8是Unicode的存储实现（转换为字节序列的规则）

	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		// %[1]x
		t.Logf("%[1]c %[1]d", c)
	}
}
