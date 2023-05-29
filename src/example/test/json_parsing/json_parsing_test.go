package json_parsing

import (
	"encoding/json"
	"testing"
)

// json字符串转换为结构
func TestJsonParsing(t *testing.T) {
	var jsonStr = `{"name":"张三","age":18}`
	var person Person
	//第一个参数必须是字节切片类型，所以需要通过 []byte 函数将字符串转换为字节切片
	//第二个参数是解析后的结果所要赋值的变量的指针
	err := json.Unmarshal([]byte(jsonStr), &person)
	if err != nil {
		t.Error("json parsing error :", err)
		return
	}
	t.Log("person :", person)
}

// json字符串转换为map
func TestJson2Map(t *testing.T) {
	var jsonStr = `{"name":"张三","age":18}`
	var data map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		t.Error("json parsing error :", err)
		return
	}
	t.Log("data :", data)
	t.Log("data Name:", data["name"])
	t.Log("data Age:", data["age"])
}

// json数组转换为切片
func TestJson2Slice(t *testing.T) {
	var jsonStr = `[{"name":"张三","age":18},{"name":"李四","age":20}]`
	var persons []Person

	err := json.Unmarshal([]byte(jsonStr), &persons)
	if err != nil {
		t.Error("json parsing error :", err)
		return
	}
	t.Log("data :", persons)
}
