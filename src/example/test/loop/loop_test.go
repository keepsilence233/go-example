package loop_test

import "testing"

// go 循环
func TestLoop(t *testing.T) {

	//条件循环
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}

	//无限循环 this loop will spin, using 100% CPU (SA5002)
	for {
		//...
	}

}
