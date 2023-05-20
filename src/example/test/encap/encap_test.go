package encap_test

import (
	"fmt"
	"testing"
	"unsafe"
)

// 结构定义
type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"1", "zhangsan", 20}
	t.Log(e)

	e1 := Employee{Name: "lisi", Age: 23}
	t.Log(e1)

	e2 := new(Employee)
	e2.Id = "2"
	e2.Name = "wangwu"
	e2.Age = 20
	t.Log(e2)
}

//行为(方法)定义

//该方式在实例对应方法调用时，实例的成员会进行值复制
// func (e Employee) toString() string {
// 	fmt.Printf("e address is %x \n", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID:%s Name:%s Age:%d", e.Id, e.Name, e.Age)
// }

// 通常情况下为了避免内存拷贝使用该方式
// toString:方法名称 string:返回类型
func (e *Employee) toString() string {
	fmt.Printf("e address is %x \n", unsafe.Pointer(&e.Name)) //获取指针地址
	return fmt.Sprintf("ID:%s Name:%s Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"1", "zhangsan", 20}
	fmt.Printf("e address is %x \n", unsafe.Pointer(&e.Name))
	t.Log(e.toString())
}
