package ch5

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}

	// 无线循环
	/*n = 0
	for {
		t.Log(n)
	}*/
}
