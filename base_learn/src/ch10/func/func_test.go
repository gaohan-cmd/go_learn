package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {

	return rand.Intn(10), rand.Intn(10)
}

func timeSpent(inner func(ap int) int) func(op int) int {
	// 计算运行某个函数的时间 传入函数 返回函数（这里的函数要返回耗时）
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Println("timeSpent", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

// 可变参数
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4, 5))
	t.Log(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func clear() {
	fmt.Println("clear resources")

}

// defer 延迟执行
// 这里可以借鉴java中的try-catch-finally的写法
func TestDefer(t *testing.T) {
	// defer 会在函数执行完毕后执行,即在函数返回时候返回，即使panic了也会执行
	// 清理某些资源 关闭某些锁
	defer func() {
		t.Log("clear resources")
	}()
	t.Log("start")
	panic("err")
	//t.Log("end")
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	// 这里类似java中装饰器模式，当某个函数经过timeSpent装饰包装后，那么它的函数功能则进行了扩展
	//(返回的函数是装饰后的函数，其具备的功能则是通过传入参数即函数作为装饰品进行装饰）
	// 这里的slowFun函数就是被装饰的函数，timeSpent函数就是装饰器，装饰器函数返回的函数就是扩展后的函数
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}
