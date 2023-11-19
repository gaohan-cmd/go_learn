package util_anyone_reply

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	//ch := make(chan string)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			//从没有消息到第一个人发了消息，就会从阻塞当中被唤醒
			//不会被阻塞，程序会继续执行
			ch <- ret

			//ch := make(chan int, 1)如果 channel 是有缓冲的（buffered channel），则在写入消息直到缓冲区满之前，其他消息可以继续进来；
			//如果 channel 是无缓冲的（unbuffered channel），则在接收者接收消息之前，发送者会被阻塞，直到消息被接收
			//当第一个消息被接收后，其他人发送消息就没有人接收消息
			//根据channel的定义，这个时候程序会被阻塞在这里
		}(i)
	}
	//一旦获取到数据，函数就会立即返回
	//有buffer的时候，不需要等待有有消息的接受者把消息拿走
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
	//开始有两条，结果结束后本来应该销毁所有的协程
	//但是反而有了11条协程，每次调用都有大量的无关进程被阻塞，会导致系统资源的耗尽
	//不断的有协程被阻塞，协程泄露，系统资源耗尽 buffer.channel放些内存泄漏
}
