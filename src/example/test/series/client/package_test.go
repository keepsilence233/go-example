package client_test

import (
	"fmt"
	"sample-app/src/example/test/series"
	"testing"
)

func init() {
	fmt.Println("clien init method ...")
}

func TestPackage(t *testing.T) {
	t.Log(series.PrintInt(1))
	t.Log(series.PrintString("1"))
}
