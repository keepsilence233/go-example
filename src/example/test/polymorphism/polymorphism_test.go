package dt_test

import (
	"fmt"
	"testing"
)

//go 多态

//自定义类型
type Code string

//接口
type Programmer interface{
	WriteHelloWorld() Code
}

//实现
type GoProgrammer struct{}
type JavaProgrammer struct{}

//重写方法
func (p *GoProgrammer) WriteHelloWorld() Code{
	return "GoProgrammer write Go HelloWorld!";
}
func (p *JavaProgrammer) WriteHelloWorld() Code{
	return "JavaProgrammer write Java HelloWorld!"
}

func WriteHelloWorld(p Programmer){
	fmt.Println(p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T){
	gp := new(GoProgrammer)
	WriteHelloWorld(gp)
	jp := new(JavaProgrammer)
	WriteHelloWorld(jp)
}