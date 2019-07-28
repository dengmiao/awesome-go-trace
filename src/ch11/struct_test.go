package ch11

import (
	"fmt"
	"testing"
	"unsafe"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u *User) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&u.Name))
	return fmt.Sprintf("Id: %d, Name: %x, Age: %d", u.Id, u.Name, u.Age)
}

/*func (u User) String() string  {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&u.Name))
	return fmt.Sprintf("Id: %d, Name: %x, Age: %d", u.Id, u.Name, u.Age)
}*/

func TestInstance(t *testing.T) {
	u1 := User{1, "Bob", 20}
	t.Log(u1)
	u2 := User{Id: 2, Age: 18, Name: "Tom"}
	t.Log(u2)
	// 返回的引用/指针 相当于 u := &User{}
	u3 := new(User)
	// 通过实例指针访问成员不需要使用->
	u3.Id = 3
	t.Logf("u2 is %T", u2)
	t.Logf("&u2 is %T", &u2)
	t.Logf("u3 is %T", u3)
}

func TestStructureOperations(t *testing.T) {
	u := User{1, "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&u.Name))
	t.Log(u.String())
}
