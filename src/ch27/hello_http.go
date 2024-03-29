package main

import (
	"fmt"
	"net/http"
	"time"
)

// 路由规则
// URL分为两种 末尾是/：表示一个子树 后面可以跟其他子路径 末尾不是/： 表示一个叶子 固定的路径
// 以/结尾的URL可以匹配它的任何子路径 比如/images会匹配/images/xx.jpg
// 它采用最长匹配原则 如果有多个匹配 一定采用匹配路径最长的那个进行处理
// 如果没有找到任何匹配项 会返回404
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.HandleFunc("/time/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t)
		w.Write([]byte(timeStr))
	})

	_ = http.ListenAndServe(":8080", nil)
}
