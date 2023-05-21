package error_test

import (
	"errors"
	"testing"
)

var ErrLessThanTwoError = errors.New("n should be in not less than 2")
var ErrLargerThenHundredError = errors.New("n should be not larger than 100")

func print(n int) (int, error) {
	if n < 2 {
		return 0, ErrLessThanTwoError
	}
	if n > 100 {
		return 0, ErrLargerThenHundredError
	}
	return n + 1, nil
}

func Test(t *testing.T) {
	//错误检查
	if v, err := print(-1); err != nil {
		if err == ErrLessThanTwoError {
			t.Log("error is LessThanTwoError:", err)
		} else if err == ErrLargerThenHundredError {
			t.Log("error is LargerThenHundredError:", err)
		} else {
			t.Error(err)
		}
	} else {
		t.Log(v)
	}
}
