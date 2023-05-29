package json_parsing

import (
	"encoding/json"
	"github.com/mailru/easyjson"
	"testing"
)

// 编码
func TestPerson_MarshalEasyJSON(t *testing.T) {
	person := Person{Name: "Tom", Age: 23}
	jsonData, err := json.Marshal(person)
	if err != nil {
		t.Error("json parsing error ...")
		return
	}
	t.Log("json str :", string(jsonData))
}

// 解码
func TestPerson_UnmarshalEasyJSON(t *testing.T) {
	jsonData := []byte(`{"name":"Tom","age":25}`)
	var person Person
	err := easyjson.Unmarshal(jsonData, &person)
	if err != nil {
		t.Error("json parsing error ...")
		return
	}
	t.Log("person :", person)
}
