package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) { //使用 go 关键字将函数作为 goroutine 启动
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second * 1)
}
