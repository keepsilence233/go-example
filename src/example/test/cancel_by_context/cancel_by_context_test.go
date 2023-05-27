package cancel_by_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/*
*
通过context判断channel是否关闭
*/
func isCancelled(context context.Context) bool {
	select {
	case <-context.Done(): //接收消息通知
		return true
	default:
		return false
	}
}

/*
*
任务取消
*/
func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background()) //创建context
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled")
		}(i, ctx)
	}
	cancel() //取消context
	time.Sleep(time.Second * 1)
}
