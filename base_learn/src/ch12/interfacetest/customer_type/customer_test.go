package customer_type

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// 定义一个函数类型,自定义类型
type IntConv func(op int) string

func timeSpent(inner func(ap int) int) IntConv {
	// 计算运行某个函数的时间 传入函数 返回函数（这里的函数要返回耗时）
	return func(op int) string {
		start := time.Now()
		// 这里只是相当于执行了传入的函数
		ret := inner(op)
		fmt.Println("timeSpent", time.Since(start).Seconds())
		return strconv.Itoa(ret)
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {

	// 这里类似java中装饰器模式，当某个函数经过timeSpent装饰包装后，那么它的函数功能则进行了扩展
	//(返回的函数是装饰后的函数，其具备的功能则是通过传入参数即函数作为装饰品进行装饰）
	// 这里的slowFun函数就是被装饰的函数，timeSpent函数就是装饰器，装饰器函数返回的函数就是扩展后的函数
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}
