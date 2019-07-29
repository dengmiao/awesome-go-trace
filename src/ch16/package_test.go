package ch16

// 没加入gopath src是gopath的一部分
import (
	"awesome-go-trace/src/ch16/series"
	"testing"
)

func TestPackage(t *testing.T) {
	// 不能访问在包外的小写方法
	t.Log(series.GetFibonacci(5))
	t.Log(series.Square(5))
}
