package unit_test

import (
	unit "sample-app/src/example/test/unit_test"
	"testing"
)

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}

	for i := 0; i < len(inputs); i++ {
		ret := unit.Square(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d,the expected id %d,the actual %d", inputs[i], expected[i], ret)
		}
	}

}
