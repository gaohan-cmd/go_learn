package csp

import (
	"fmt"
	"testing"
	"time"
)

// 模拟了一个网络请求的方法调用过程，通过Channel来控制当前协程在网络请求的等待过程中，去执行别的任务

func serviceTask() string {
	fmt.Println("--start serviceTask--")
	time.Sleep(time.Millisecond * 50)
	return "--serviceTask Done--"
}

// 别的任务
func otherTask() {
	fmt.Println("--start otherTask--")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("--otherTask Done--")
}

// csp异步管道 对上述serviceTask方法进行包装
func AsyncService() chan string {
	// 当不加入这个限定容量时候 那么协程没有被调用时候，会一直阻塞在这里，通过限定容量，可以让协程不阻塞，因为已知消息只有一个
	retCh := make(chan string)
	// 创建一个channel 缓冲大小为 1 的字符串类型通道
	// retCh := make(chan string, 1)
	// 开启一个协程
	go func() {
		// 调用serviceTask方法 将结果返回给ret，下面把ret放入到channel里 当其他程序需要结果时候 可以在channel上进行等待
		ret := serviceTask()
		fmt.Println("returned result.")
		// 将结果放入管道中
		retCh <- ret
		fmt.Println("service exited.")
	}()
	// 返回管道
	return retCh
}

func TestAsyncService(t *testing.T) {
	// 调用异步管道
	retCh := AsyncService()
	// 调用别的任务
	otherTask()
	// 从管道中读取结果，如果不按上面限定容量的话 那么会一直阻塞直到这个时候才能被释放即打印
	fmt.Println(<-retCh)
	// 等待异步管道执行完毕
	time.Sleep(time.Second * 1)

}

// 不进行并发管理 串行输出运行
func TestAsyncServiceWithNo(t *testing.T) {

	// 从管道中读取结果
	fmt.Println(serviceTask())
	// 调用别的任务
	otherTask()

}
