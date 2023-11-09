package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

/**
1 << 0 运算是位运算之一，表示将数字 1 左移 0 位。在二进制中，
左移相当于向左移动数字中的每一位，并在右侧填充零。因为这里是左移 0 位，所以没有进行任何位移，结果仍然是 1。

更具体的解释：
数字 1 的二进制表示是 00000001。
左移 0 位后，数字不会发生位移，仍然是 00000001。由于左侧没有位移，所以可以认为左移操作相当于乘以 2 的幂次方，
其中幂次方等于位移的位数。因此，1 << 0 的结果等于 1 * 2^0，即 1。
因此，1 << 0 等于 1。同样的，如果将数字 1 左移 1 位，结果将是 2，因为移位后的二进制为 00000010，相当于数值 2 的二进制表示。
*/
// 对于连续位的常量可以使用以下写法
const (
	ReadAble = 1 << iota
	WriteAble
	ExclusiveAble
)

func TestConstantTry1(t *testing.T) {
	a := 7 //0111
	t.Log(ReadAble, WriteAble, ExclusiveAble)
	t.Log(a&ReadAble == ReadAble, a&WriteAble == WriteAble, a&ExclusiveAble == ExclusiveAble)

}

func TestConstantTry(t *testing.T) {
	t.Log("星期六", Saturday)
	t.Log("星期日", Sunday)
}
