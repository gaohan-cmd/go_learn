package fib

import (
	"testing"
)

// 方法推荐使用驼峰命名 小写私有 大写共有
func TestFibList(t *testing.T) {
	var a int = 1
	var b int = 1
	t.Log(a, "")
	for i := 0; i < 5; i++ {
		t.Log(b, " ")
		temp := a
		a = b
		b = temp + a
	}
	t.Log(a, "")
}

func TestChange(t *testing.T) {
	a := 1
	b := 2
	temp := b
	b = a
	a = temp
	t.Log("初始 ab", a, b)
	t.Log("后来 ab", a, b)

	a, b = b, a
	t.Log("简化版交换值", a, b)
}
