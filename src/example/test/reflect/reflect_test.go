package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

// CheckType 通过kind来判断类型
func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)
	}
}
func TestBasicType(t *testing.T) {
	var f float32 = 1
	CheckType(f)
}

type Employee struct {
	EmployeeId string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newAge int) {
	e.Age = newAge
}

type Customer struct {
	CookieId string
	Name     string
	Age      int
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Tom", 30}
	//按名字获取成员
	t.Log(reflect.ValueOf(*e).FieldByName(`Name`))
	t.Logf("Name:value(%[1]v),Type((%[1]T)", reflect.ValueOf(*e).FieldByName(`Name`))

	if nameField, ok := reflect.TypeOf(*e).FieldByName(`Name`); !ok {
		t.Error("Failed to get Name field.")
	} else {
		t.Log("Tag:format:", nameField.Tag.Get("format"))
	}

	//动态调用method
	reflect.ValueOf(e).MethodByName("UpdateAge").
		Call([]reflect.Value{reflect.ValueOf(23)})
	t.Log("UpdateAge:", e)
}
