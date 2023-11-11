package series

// 在main被执行前，所有依赖的package的init 方法都会被执行
// 不同包的init函数按照包导入的依赖关系决定执行顺序
// 每个包可以有多个init函数.
// 包的每个源文件也可以有多个init函数，这点比较特殊.
func init() {
	println("init1")
}

func init() {
	println("init2")
}

func GetFibonacci(n int) []int {
	fibonacciList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibonacciList = append(fibonacciList, fibonacciList[i-1]+fibonacciList[i-2])
	}
	return fibonacciList
}

// 小写的方法在包外不能被访问
func square(n int) int {
	return n * n
}
