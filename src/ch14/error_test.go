package ch14

import (
	"github.com/pkg/errors"
	"testing"
)

var LessThanTowError = errors.New("n cloud not less than 2")
var LargerThanHundredError = errors.New("n cloud not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTowError
	}
	if n > 100 {
		return nil, LargerThanHundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(100); err != nil {
		switch err {
		case LessThanTowError:
			t.Error("too less")
		case LargerThanHundredError:
			t.Error("too larger")
		default:
			t.Error(err)
		}
	} else {
		t.Log(v)
	}
}
