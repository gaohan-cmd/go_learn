package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

// 结构体绑定定义
func (e Employee) String() string {
	// 调用fmt.Sprintf函数，可以将这些字段的值按照指定的格式组合成一个字符串，并作为方法的返回值
	return fmt.Sprintf("Id:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)

}

// 指针绑定定义
func (e *Employee) String1() string {
	// 调用fmt.Sprintf函数，可以将这些字段的值按照指定的格式组合成一个字符串，并作为方法的返回值
	return fmt.Sprintf("Id:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)

}

// 指针绑定定义
func (e *Employee) String2() string {
	// 不安全编程
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)

}

// 结构体绑定定义
func (e Employee) String3() string {
	// 调用fmt.Sprintf函数，可以将这些字段的值按照指定的格式组合成一个字符串，并作为方法的返回值
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)

}

func TestCreateEmployeeObj(t *testing.T) {

	// 方法一
	e := Employee{
		"0", "Bob", 20}
	// 方法二
	e1 := Employee{
		Name: "Mike", Age: 30}
	// 方法三，使用new关键字(返回的是指针，使用.访问数据)
	e2 := new(Employee) // 返回指针
	e2.Id = "2"
	e2.Name = "Rose"
	e2.Age = 22
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)
}

/*
*
绑定在结构体上的方法
此demo演示结构体和指针均能调用被绑定的方法
*/
func TestStructOprations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	// 当定义为结构体时候作为结构体会被调用
	t.Log(e.String())
	// 当作为指针时候 也能访问在结构上定义的方法
	e1 := &Employee{"0", "Bob", 20}
	// 当定义为结构体时候作为结构体会被调用
	t.Log(e1.String())

}

/*
*
绑定在指针上的方法
此demo演示结构体和指针均能调用被绑定的方法
*/
func TestStructOprations1(t *testing.T) {
	// 当定义为指针时候 结构体和指针的方式也均能调用该方法
	e := Employee{"0", "Bob", 20}

	t.Log(e.String1())

	e1 := &Employee{"0", "Bob", 20}

	t.Log(e1.String1())

}

/*
*
绑定在指针上的方法
此demo演示绑定指针不会出现对象复制 减小内存开销
*/
func TestStructOprations2(t *testing.T) {
	// 当定义为指针时候 结构体和指针的方式也均能调用该方法
	e := Employee{"0", "Bob", 20}
	// 当定义为结构体，不会出现对象复制，
	// 但是这里需要注意调用时候需要使用结构体去调用，如果定义指针去调用，那么化石会去对象复制
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))

	t.Log(e.String2())

}

/*
*
绑定在指针上的方法
此demo演示绑定结构体会出现对象复制
*/
func TestStructOprations3(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))

	e2 := new(Employee)
	e2.Id = "2"
	e2.Name = "Rose"
	// 当结构体绑定定义时候 结构体里的数据被复制了一份
	t.Log(e2.String2())

}
