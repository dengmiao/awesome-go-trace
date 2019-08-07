package ch24

import (
	"reflect"
	"testing"
)

// 按名字访问结构成员
// reflect.ValueOf(*e).FieldByName("name")

// 按名字访问结构的方法
// reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(age int) {
	e.Age = age
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 20}
	// 按名字获取成员
	t.Logf("Name: value(%[1]v), Type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'name' field.")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Update Age:", e)
}
