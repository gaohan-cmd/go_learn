package slic_test

import "testing"

func TestSliceInit(t *testing.T) {

	// 这里没有写具体的长度 是可变长的
	var s0 []int
	t.Log(len(s0), cap(s0))
	// 切片内部结构 ptr(指针 指向连续的存储空间 数组）、len(元素的个数） cap(内部数组的容量）
	// 通过使用append()函数将值为1的元素追加到切片s0的末尾。这会导致切片s0的长度增加为1，并且它的容量也会自动调整
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	// 使用make创建不同类型
	/**
	make()函数在Go语言中主要用于创建slice、map和channel等类型的对象，它们都是引用类型。
	创建slice类型
	对于slice类型来说，make()函数需要两个参数：类型和长度。其中，类型为slice，长度表示要分配的底层数组长度，容量默认等于长度。

	例如，下面的代码创建了一个长度为5的int类型切片：

	s := make([]int, 5)
	还可以通过传递第三个参数来指定切片的容量：

	s := make([]int, 5, 10)
	这里表示创建一个长度为5、容量为10的int类型切片。

	创建map类型
	对于map类型来说，make()函数只需要一个参数：map类型。并且返回一个空的、可用于存储键值对的map变量。

	例如，下面的代码创建了一个string类型到int类型的映射map：

	m := make(map[string]int)
	创建channel类型
	对于channel类型来说，make()函数需要一个参数：channel类型。并且返回一个空的、可用于通信的channel变量。

	例如，下面的代码创建了一个int类型的通道：

	c := make(chan int)
	当然，还可以通过传递第二个参数来指定缓冲区大小，例如：

	c := make(chan int, 10)
	这里表示创建一个容量为10的带缓冲区的int类型通道，可以存储最多10个元素。

	例子：
	make(map[string]chan []byte, 80) 可以看出，这里创建了一个映射，键为 string 类型，值为 chan []byte 类型
	*/
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	// len长度是多少就初始化多少 ,但如果继续append的话那么可以继续访问，
	//cap代表容量 如果比如append(s2, 4)超过容量那么就会自动扩容 在实际开发过程会提前设定cap以防多次扩容影响效率
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	s2 = append(s2, 1)
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3], s2[4], s2[5])

}

func TestSliceGrowing(t *testing.T) {

	// 使用自动识别类型时候 需要添加{}
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}

}

func TestSliceMemory(t *testing.T) {
	var s []int
	t.Log(s)
	year := []string{"JaN", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	//slice本身指向就是数组 不同slice对同一个数组进行切分 那么相当于共享一个数组
	// 下面Q2的cap是9 因为从索引3切分后剩下还有9个
	t.Log(Q2, len(Q2), cap(Q2))

	summary := year[5:8]
	t.Log(summary, len(summary), cap(summary))

	// 在这里如果进行修改 那么其他slice里的数值也会修改 因为是共享的
	summary[0] = "Unkown"
	t.Log(Q2)
}

func TestSliceCpmpare(t *testing.T) {
	// slice不能用作比较 只能和nil进行比较
	/*	a := []int{1, 2, 3, 4}
		b := []int{1, 2, 3, 4}
		a = b
			if (a == b) {
			t.Log("equal")
		}*/

}
