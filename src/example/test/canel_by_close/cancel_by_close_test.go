package canel_by_close

import (
	"fmt"
	"testing"
	"time"
)

/*
*
判断任务是否取消
*/
func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

//关闭channel
func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}
func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

/*
*
任务取消
*/
func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled")
		}(i, cancelChan)
	}
	cancel_2(cancelChan)
	time.Sleep(time.Second * 1)
}
