package ch12

import (
	"fmt"
	"testing"
)

type pet struct {
}

func (p *pet) Speak() {
	fmt.Print("...\n")
}

func (p *pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println("", host+"\n")
}

type Dog struct {
	pet
}

func (d *Dog) SpeakTo(host string) {
	fmt.Println("wang!", host)
}

func TestExtension(t *testing.T) {
	dog := new(Dog)
	dog.Speak()
	dog.SpeakTo("Tom")

	// 不支持继承 不支持显示类型转化
	//var dog2 pet = new(Dog)
}
