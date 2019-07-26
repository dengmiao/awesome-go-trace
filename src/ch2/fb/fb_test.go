package fb

import (
	"fmt"
	"testing"
)

/**
斐波那契
*/
func TestFb(t *testing.T) {
	// go中申明变量 必须使用
	var a int = 1
	// 申明赋值的另一种形式
	b := 1
	// 默认类型可省略
	/*
		var (
			a int = 1
			b = 1
		)
	*/
	fmt.Print(a)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}

/**
交换变量
*/
func TestExchange(t *testing.T) {
	a := 2
	b := 9
	tmp := a
	a = b
	b = tmp
	t.Log(a, b)
}
