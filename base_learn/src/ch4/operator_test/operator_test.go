package operator_test

import "testing"

const (
	ReadAble = 1 << iota
	WriteAble
	ExclusiveAble
)

func TestCompareArray(t *testing.T) {
	// 定长数组的声明方式，即使用 ... 省略数组长度
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 4, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == b)
	// 长度不同的会报编译错误
	//t.Log(a==c)
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	a := 7 //0111
	t.Log(ReadAble, WriteAble, ExclusiveAble)
	// a &^ b，表示将 a 中和 b 对应二进制位都是 1 的位清零
	a = a &^ ReadAble
	t.Log(a&ReadAble == ReadAble, a&WriteAble == WriteAble, a&ExclusiveAble == ExclusiveAble)

}
