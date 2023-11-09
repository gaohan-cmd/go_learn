package type_test

import (
	"testing"
)

// 自定义别名
type MyInt int64

func TestImplict(t *testing.T) {
	// 不支持隐式类型转换
	var a int32 = 1
	var b int64
	//b = a
	b = int64(a)
	t.Log(a, b)
	// 别名隐式类型也不允许
	var c MyInt
	//c = b
	c = MyInt(b)
	t.Log(a, c, c)

}

func TestPont(t *testing.T) {
	a := 1
	aPtr := &a
	// Go语言中不支持指针计算
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)

}

func TestString(t *testing.T) {
	// s初始化时候 被初始化为空字符串
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))

	if s == "" {
		t.Log("s is empty")
	}
}
