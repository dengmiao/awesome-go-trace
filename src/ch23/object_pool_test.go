package ch23

import (
	"fmt"
	"testing"
	"time"
)

func TestObjectPool(t *testing.T) {
	pool := NewObjPool(10)
	if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
		t.Error(err)
	}
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}

	}

	fmt.Println("Done")
}
