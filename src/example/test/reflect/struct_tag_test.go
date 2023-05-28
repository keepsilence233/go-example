package reflect

import (
	"reflect"
	"testing"
)

type Person struct {
	Name  string `json:"name" xml:"name"`
	Age   int    `json:"age" xml:"age"`
	Phone string `json:"phone" xml:"phone"`
}

// TestStructTag 访问 struct tag
func TestStructTag(t *testing.T) {
	p := &Person{"Tom", 23, "6666666"}

	if field, ok := reflect.TypeOf(*p).FieldByName("Name"); ok {
		t.Log("field tag json:", field.Tag.Get("json"))
		t.Log("field tag xml:", field.Tag.Get("xml"))
	}
}
