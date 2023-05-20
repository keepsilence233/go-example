package interface_test

import "testing"

// 接口定义
type Programmer interface {
	WriteHelloWorld() string
}

// 接口实现
type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() string {
	return "HelloWrold!"
}

func TestClient(t *testing.T) {
	var c Programmer = new(GoProgrammer)
	i := GoProgrammer{}
	t.Log(c.WriteHelloWorld())
	t.Log(i.WriteHelloWorld())
}
