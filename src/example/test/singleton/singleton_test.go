package singleton

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singletonInstance *Singleton
var once sync.Once

// 创建一个单例的对象
func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("create obj")
		singletonInstance = new(Singleton)
	})
	return singletonInstance
}

func TestSingleton(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%d\n", unsafe.Pointer(obj)) //输出单例对象的地址
			wg.Done()
		}()
	}
	wg.Wait()
}
