package opertaortest_test

import "testing"

// go 运算符
func TestOperator(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 4, 5}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == b)
	// t.Log(a == c)  //不同长度的数组比较会编译错误	
	t.Log(a == d)

	ia := 1
	ib := 2
	t.Log(ia + ib)
	t.Log(ia - ib)
	t.Log(ia * ib)
	t.Log(ia / ib)
}
