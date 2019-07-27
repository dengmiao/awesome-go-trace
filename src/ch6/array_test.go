package ch6

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	// 初始化为0值
	t.Log(arr[0])
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 4, 5}
	t.Log(arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr2 := [...]int{1, 2, 4, 5}
	for i := 0; i < len(arr2); i++ {
		t.Log(arr2[i])
	}

	// index, item := range arr , 不关系index 用_占位
	for _, e := range arr2 {
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr2 := [...]int{1, 2, 3, 4, 5}
	// 前三个
	arr2_sec0 := arr2[:3]
	// 下标为3后的元素 :前后不填写 则获取所有元素
	arr2_sec1 := arr2[3:]
	t.Log(arr2_sec0, arr2_sec1)
}
