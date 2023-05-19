package consttest_test

import (
	"testing"
)

/*
常量
iota:连续常量定义
*/
const (
	Mondy = 1 + iota	
	Tuesday
	Wednesday
)

func TestConst(t *testing.T) {
	t.Log(Mondy, Tuesday, Wednesday)
}
