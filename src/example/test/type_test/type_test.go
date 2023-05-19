package try_test

import "testing"

type myInt int64 //别名

//类型转换
func TestImplicit(t *testing.T) {
	// var a int32 = 1
	// var b int64 = 2
	// b = a 	//go 不支持隐式类型转换
	// t.Log(a,b)

	var a int32 = 1
	var b int64
	b = int64(a)
	var c myInt
	c = myInt(b)
	t.Log(a, b, c)
}

//指针
func TestPoint(t *testing.T){
	a:=1
	aPar := &a
	t.Log(a,aPar)
	t.Logf("%T %T",a,aPar)
}

//字符串
func TestString(t *testing.T){
	var s string 	//会被初始化为空字符串
	t.Log("*"+s+"*")
	t.Log(len(s))
}