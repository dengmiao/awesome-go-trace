package main

import (
	"fmt"
	"os"
)

// 主函数必须在main包下 可以不是main目录 main函数不支持入参、出参
func main() {
	if len(os.Args) > 0 {
		fmt.Println("Hello World", os.Args[0])
	}
	// 0
	os.Exit(-1)
}
