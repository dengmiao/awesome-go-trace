package ch5

import "testing"

func TestIfMultiSec(t *testing.T) {
	// if var declaration; condition {}
	// condition 为表达式结果 boolean
	if a := 1 == 1; a {
		t.Log("a == 1", a)
	}
}

// switch 条件表达式不限制常量或整数
// 不需要case中加break
// case中可以出现多个结果项(命中任意一个即可执行) 用逗号分隔
// 可以不设定switch后的条件表达式 这种情况下 和多个if .. else if ... else ..逻辑相同
func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log(i, "Even")
		case 1, 3:
			t.Log(i, "Odd")
		default:
			t.Log(i, "not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log(i, "Even")
		case i%2 == 1:
			t.Log(i, "Odd")
		default:
			t.Log("unknown")
		}
	}
}
