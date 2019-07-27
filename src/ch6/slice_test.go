package ch6

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	// len 可访问/初始化元素个数  cap 容量
	t.Log(s0, len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(s0, len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(s1, len(s1), cap(s1))

	// 未初始化元素不可访问
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	// 扩容 *2增长
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	mon := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := mon[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := mon[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "unknown"
	t.Log(Q2)
}

// 数组 vs 切片
// 容量是否可伸缩 是否可比较
// 数组可比较不可扩容 切片可扩容不可比较
func TestSliceCompare(t *testing.T) {
	//s1 := []int{1, 2, 3}
	//s2 := []int{1, 2, 4}
	//t.Log(s1 == s2)
}
