package ch24

import (
	"reflect"
	"testing"
)

// map和切片 比较
func TestDeepEqual(t *testing.T) {
	a := map[string]int{"a": 1, "b": 2, "c": 3}
	b := map[string]int{"a": 1, "b": 2, "c": 3}
	t.Log("a == b", reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 3, 2}
	s3 := []int{1, 2, 3}
	t.Log("s1 == s2", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3", reflect.DeepEqual(s1, s3))
}

func TestFillNameAndAge(t *testing.T) {

}
