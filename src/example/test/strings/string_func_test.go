package string_test

import (
	"strconv"
	"strings"
	"testing"
)

// go string func
func TestStringFunc(t *testing.T) {
	s := "A,B,C"

	parts := strings.Split(s, ",")
	t.Log(parts)
	for _, item := range parts {
		t.Log(item)
	}

	t.Log(strings.Join(parts,"-"))
}

//类型转换
func TestConv(t *testing.T){
	s := strconv.Itoa(10)
	t.Log("str :"+s)

	//t.Log(10 + strconv.Atoi(s))

	if i,err:=strconv.Atoi(s);err==nil{
		t.Log(10 + i)
	}
}
