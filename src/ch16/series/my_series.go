package series

import (
	"errors"
	"fmt"
)

var LessThanTowError = errors.New("n cloud not less than 2")
var LargerThanHundredError = errors.New("n cloud not larger than 100")

// 小写方法不能在包外被访问
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

// main方法执行前 所有依赖的package的init方法都会被执行
// 不同包的init函数按照导入的依赖关系决定执行顺序
// 每个包可以有多个init函数
// 包的每个源文件也可以有多个init函数
func init() {
	fmt.Println("init1...")
}

func init() {
	fmt.Println("init2...")
}

func Square(n int) int {
	return n * n
}
