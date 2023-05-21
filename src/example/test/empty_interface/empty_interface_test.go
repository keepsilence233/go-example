package emptyinterface_test

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	/*
		ok := p.(int) : 断言,如果被断言为一个整型
		v : 转换后的值
	*/
	if v, ok := p.(int); ok {
		fmt.Println("Integer", v)
		return
	}
	if v, ok := p.(string); ok {
		fmt.Println("string", v)
		return
	}
	fmt.Println("unknow type")
}

func DoSomethingWithSwitch(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("unknow type")
	}
}

func TestEmptyInterface(t *testing.T) {
	DoSomethingWithSwitch(1)
	DoSomethingWithSwitch("666")
	DoSomethingWithSwitch(t)
}
