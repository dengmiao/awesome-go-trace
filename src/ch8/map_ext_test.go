package ch8

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 2
	isExisting := isExist(n, mySet)
	t.Logf("%d is existing or not ? %v", n, isExisting)
	mySet[2] = true
	t.Log(len(mySet))
	delete(mySet, 1)
	n = 1
	isExisting = isExist(n, mySet)
	t.Logf("%d is existing or not ? %v", n, isExisting)
}

func isExist(key int, mySet map[int]bool) bool {
	if mySet[key] {
		return true
	} else {
		return false
	}
}
