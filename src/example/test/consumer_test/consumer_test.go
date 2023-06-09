package consumertest_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//自定义类型
type IntConv func (op int) int

//使用自定义类型
func timeSpent(inner func(op int) int) IntConv{
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

func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}
