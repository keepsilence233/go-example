package obj_pool

import (
	"fmt"
	"testing"
	"time"
)

// 测试对象池
func TestConnectionPool(t *testing.T) {
	pool := NewConnectionPool(10)
	for i := 0; i < 11; i++ {
		if v, err := pool.GetConnection(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			if err := pool.ReleaseConnection(v); err != nil {
				t.Error(err)
			}
		}
	}
	t.Log("Done")
}
