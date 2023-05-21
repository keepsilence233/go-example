package remote_package_test

import (
	"testing"
	//引入远程包,可以取一个cm的别名
	cm "github.com/easierway/concurrent_map"
)

func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(32)
	m.Set(cm.StrKey("one"), 1)
	t.Log(m.Get(cm.StrKey("one")))
	t.Log(m.Get(cm.StrKey("two")))
}
