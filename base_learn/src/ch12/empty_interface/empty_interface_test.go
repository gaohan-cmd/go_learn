package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) { // 传入的是个空接口
	// 如果传入参数能够被断言成一个整形
	if i, ok := p.(int); ok { // 断言判断
		fmt.Println("Integer: ", i)
		return
	}
	if s, ok := p.(string); ok {
		fmt.Println("string: ", s)
		return
	}
	fmt.Println("Unknow Type.")
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}

// 以上方法有点繁琐，最好使用switch结构
func DoSomething2(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer: ", v)
	case string:
		fmt.Println("string: ", v)
	default:
		fmt.Println("Unknow Type.")
	}
}

func TestEmptyInterfaceAssertion2(t *testing.T) {
	DoSomething2(10)
	DoSomething2("10")
}
