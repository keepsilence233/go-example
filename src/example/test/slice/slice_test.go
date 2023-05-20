package slice_test

import "testing"

// go 切片
func TestSliceInit(t *testing.T) {

	var s0 []int
	t.Log(s0, len(s0), cap(s0))

	s0 = append(s0, 1) //添加参数
	t.Log(s0, len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(s1, len(s1), cap(s1))

	s2 := make([]int, 3, 5) //长度为3 容量为5
	t.Log(s2, len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 666)
	t.Log(s2[0], s2[1], s2[2], s2[3])
}

// go 切片增长
// go的cap增长每次是当前的二倍
func TestSliceGrowing(t *testing.T) {
	s := []int{}

	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

// 切片的共享存储
func TestSliceShareMemory(t *testing.T) {

	year := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	Q2 := year[3:6]
	t.Log(Q2,len(Q2),cap(Q2))

	summer := year[5:8]
	t.Log(summer,len(summer),cap(summer))

	t.Log(" --- --- --- ---")
	//修改之后会对所有的连续空间影响
	summer[0] = "unknow"
	t.Log(summer)
	t.Log(year)

}

//go 切片比较
func TestSliceComparing(t *testing.T){

	a := []int{1,2,3,4}
	v := []int{1,2,3,4}

	//t.Log(a == v) invalid operation: a == v (slice can only be compared to nil)
	t.Log(a == nil)
	t.Log(v == nil)
	
}