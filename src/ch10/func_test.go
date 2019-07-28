package ch10

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 函数是一等公民
// 可以有多个返回值
// 所有参数都是值传递：slice, map, channel会有传引用的错觉
// 函数可以作为变量的值
// 函数可以作为参数和返回值
func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)

	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func TestDefer(t *testing.T) {
	// defer 延迟执行
	defer Clear()
	fmt.Println("Start")
	panic("err")
	// 不可达
	//fmt.Println("End")
}

// 多返回
func returnMultiValues() (int, int) {
	return rand.Intn(20), rand.Intn(20)
}

// 计时
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

// 可变参
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func Clear() {
	fmt.Println("Clear resources.")
}
