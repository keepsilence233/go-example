package obj_cache

import (
	"fmt"
	"sync"
	"testing"
)

// TestSyncPool 测试sync.Pool
func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create new Object")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println("v:", v)

	pool.Put(666)

	//runtime.GC() //手动GC 清除sync.Pool中缓存的对象

	v1 := pool.Get().(int)
	fmt.Println("v1:", v1)
	v2 := pool.Get().(int)
	fmt.Println("v2:", v2)
}

// TestSyncPoolInMultiGroutine 测试多协程情况下sync.Pool
func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create new Object")
			return 100
		},
	}

	pool.Put(666)
	pool.Put(666)
	pool.Put(666)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(pool.Get().(int))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
