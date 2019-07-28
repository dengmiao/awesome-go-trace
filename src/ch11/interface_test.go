package ch11

import "testing"

// 接口为非入侵性的 不依赖于接口定义 duck type
// 所以接口的定义可以包含在接口使用者的包内

type Programmer interface {
	WriteHello() string
}

type GoProgrammer struct {
}

// 方法签名一致即可
func (gop *GoProgrammer) WriteHello() string {
	return "fmt.Println(\"Hello\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHello())
}
