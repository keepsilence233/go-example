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

func otherTask() {
	fmt.Println("working on something else.")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("task is done.")
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
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh) //<-retCh 从channel中取数据
}
