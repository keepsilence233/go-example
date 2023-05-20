package map_test

import "testing"

// go map init
func TestMapInit(t *testing.T) {

	m := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(m, len(m))

	m1 := map[string]int{}
	m1["one"] = 1
	m1["two"] = 2
	m1["three"] = 2
	t.Log(m1, len(m1))

	m2 := make(map[int]int, 10)
	t.Log(m2, len(m2))
}

// 访问map
func TestAccessMap(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}

	//遍历map
	for key, value := range m {
		t.Log(key, value)
	}
	t.Log(" --- --- ---")

	m1 := map[int]int{}
	//访问一个不存在的值 不存在会返回0 不是nil
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 666
	if v, ok := m1[3]; ok {
		t.Logf("key 3 value is %d", v)
	}else{
		t.Log("key 3 is not existing.")
	}
}


//value为func的map
func TestMapWithFuncValue(t *testing.T){
	m := map[int]func(op int)int{}
	m[1] = func(op int) int {return op}
	m[2] = func(op int) int {return op + 1}
	m[3] = func(op int) int {return op - 1}
	m[4] = func(op int) int {return op * 2}
	m[5] = func(op int) int {return op / 2}
	t.Log(m[1](2),m[2](2),m[3](2),m[4](2),m[5](2))
}