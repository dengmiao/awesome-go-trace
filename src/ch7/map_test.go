package ch7

import "testing"

func TestMapInit(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 4, "c": 9}
	t.Log(m1["b"])
	t.Logf("len(m1) = %d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len(m2) = %d", len(m2))
	m3 := make(map[int]int, 10)
	// cap 函数不能求map的容量
	t.Logf("len(m3) = %d", len(m3))
}

func TestAccessNotExistKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	// 访问map不存在的key 会获得零值 不能通过nil判断key是否存在
	if v, ok := m1[3]; ok {
		t.Logf("key 3's value is %d", v)
	} else {
		t.Log("key 3 is not existing")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Logf("%d's value is %d ", k, v)
	}
}
