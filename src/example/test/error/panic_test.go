package error_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	defer func() {
		fmt.Println("finally!")
	}()
	panic(errors.New("errors new!"))
}

func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil { // 在defer 函数中使用recover()函数捕获错误
			fmt.Println("finally!")
			//恢复错误...
		}
	}()
	t.Log("start")
	panic(errors.New("errors new!"))
}
