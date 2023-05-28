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
