package ch24

import (
	"errors"
	"reflect"
	"testing"
)

// map和切片 比较
func TestDeepEqual(t *testing.T) {
	a := map[string]int{"a": 1, "b": 2, "c": 3}
	b := map[string]int{"a": 1, "b": 2, "c": 3}
	t.Log("a == b", reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 3, 2}
	s3 := []int{1, 2, 3}
	t.Log("s1 == s2", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3", reflect.DeepEqual(s1, s3))
}

func TestFillNameAndAge(t *testing.T) {
	setting := map[string]interface{}{"Name": "Mike", "Age": 25}
	e := Employee{}
	if err := fillBySetting(&e, setting); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	c := new(Customer)
	if err := fillBySetting(c, setting); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}

func fillBySetting(st interface{}, setting map[string]interface{}) error {
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			return errors.New("the first param should be a pointer to the struct type")
		}
	}

	if setting == nil {
		return errors.New("setting is nil")
	}

	var (
		field reflect.StructField
		ok    bool
	)

	for k, v := range setting {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}

	return nil
}
