package csp_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func AsyncService() chan string {
	retCh := make(chan string) //声明一个channel
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret //往channel中放数据
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	//多路选择实现超时
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	// case <-time.After(time.Millisecond * 100):
	case <-time.After(time.Millisecond * 10):
		t.Error("time out")
	}
}
