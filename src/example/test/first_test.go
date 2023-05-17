package try_test

import "testing"

// go 单元测试 文件名称以_test.go结尾
func TestFirstTry(t *testing.T) {
	t.Log("My first try")
}

/*
常量
*/
const (
	Mondy = 1 + iota
	Tuesday
	Wednesday
)

func TestConstantTry(t *testing.T) {
	t.Log(Mondy, Tuesday, Wednesday)
}
