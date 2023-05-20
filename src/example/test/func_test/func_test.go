package functest_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//go多个返回值方法定义
func retrunMulitValues() (int, int) {
	return rand.Int(), rand.Int()
}

func TestFn(t *testing.T) {
	a, b := retrunMulitValues()
	t.Log(a, b)

	c,_ := retrunMulitValues()
	t.Log(c)


	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
	
}

//计算函数计算时间的公共方法
func timeSpent(inner func(op int) int) func (op int) int{
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:",time.Since(start).Seconds())
		return ret
	}
}
func slowFun(op int) int{
	time.Sleep(time.Second * 1)
	return rand.Int();
}



func TestVarParam(t *testing.T){
	t.Log(sum(1,1,1,1,1))
}
//go 可变参数
func sum(ops ...int) int {
	ret := 0
	for _,item := range ops{
		ret = ret + item
	}
	return ret;
}

func Clear(){
	fmt.Println("clear resources!")
}
//defer 函数
func TestDefer(t *testing.T){
	defer Clear()
	t.Log("start")
	t.Log("execute")
}