package string

import "testing"

func TestString(t *testing.T) {
	// string 是不可变的byte切片 不能进行s[1]这样的修改操作
	var s string
	t.Log(s)
	// var声明变量不需要{} :=可以直接自动匹配类型
	arrr := [...]int{1}
	arrr[0] = 1
	t.Log(arrr[0])

	s = "Hello"
	t.Log(len(s))
	// s[1] = ‘3’ //string 类型是不可变的slice
	s = "\xE4\xB8\xA5" // 可以存储任何二进制数据
	//s = "\xE4\xBA\xB5" // 可以存储任何二进制数据
	t.Log(s)
	// len 计算的是byte数
	t.Log(len(s))
	s = "中"
	t.Log(len(s)) // 是byte数

	// rune 能取出字符串的unicode
	c := []rune(s)
	t.Log(len(c))
	// t.Log("run size:",unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)

}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	// %d 格式化为整数 %c格式化为字符串 字符类型的值，%x 将其解释为对应的 Unicode 码点
	for _, i := range s {
		t.Logf("%[1]c %[1]x", i)
	}

}
