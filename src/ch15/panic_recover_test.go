package ch15

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExist(t *testing.T) {
	defer func() {
		//fmt.Println("Finally")
		if err := recover(); err != nil {
			fmt.Println("recover from", err)
			//t.Error(err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("something wrong"))
}
