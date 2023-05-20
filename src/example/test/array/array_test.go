package array_test

import (
	"testing"
)

// go array
func TestArrayInt(t *testing.T) {

	var array [3]int
	t.Log(array[0], array[1], array[2])
	array[0] = 0
	array[1] = 1
	array[2] = 2
	t.Log(array)

	array1 := [...]int{1, 2, 3, 4, 5}
	t.Log(array1)

}

// go 遍历数组
func TestArrayTravel(t *testing.T) {
	array := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(array); i++ {
		t.Log(array[i])
	}

	t.Log(" --- --- ---")

	for idx, e := range array {
		t.Log(idx, e)
	}
	t.Log(" --- --- ---")
	//同上,通过_代替inx来占位
	for _, e := range array {
		t.Log(e)
	}
}

// go 数组截取
func TestArraySection(t *testing.T) {

	//a[开始索引(包含),结束索引(不包含)]

	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	t.Log(a[1:2])      //2
	t.Log(a[1:3])      //2,3
	//t.Log(a[1:len(a)]) //2,3,4,5,6,7,8
	t.Log(a[1:])       //2,3,4,5,6,7,8
	t.Log(a[:3])       //1,2,3

}
