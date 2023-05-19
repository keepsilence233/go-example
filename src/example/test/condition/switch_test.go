package condition_test

import (
	"runtime"
	"testing"
)

// go switch
func TestSwitchCase(t *testing.T) {
	switch os := runtime.GOOS; os {
	case "darwin":
		t.Log("OS X.")
	case "linux":
		t.Log("Linux.")
	default:
		t.Log("%S.", os)
	}
}

func TestSwitchMulitCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unknow")
		}
	}
}
