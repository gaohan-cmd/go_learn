package cancel_by_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/*
*
●根Context:通过context. Background 0)创建
●子Context: context.WithCancel(parentContext)创建

	●ctx, cancel := context. WithCancel(context. Background()

●当前Context被取消时，基于他的子context都会被取消
●接收取消通知<-ctx.Done()
*/
func TestCancel(t *testing.T) {
	// ctx是新创建的带有取消功能的上下文对象，cancel是与之对应的取消函数
	ctx, cancel := context.WithCancel(context.Background())
	// 同时启动5个协程
	for i := 0; i < 5; i++ {
		go func(i int) {
			for {
				if isCancel(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i)
	}
	cancel()
	time.Sleep(time.Second * 1)

}

// 判断从收到的channel上是否收到了消息
func isCancel(ctx context.Context) bool {
	// ctx 表示的是上下文（Context）。上下文是一个在请求范围内携带截止日期、取消信号、请求作用域以及其他请求相关值的对象。
	//上下文一般用于在并发处理中传递请求作用域、取消信号、截止日期等与请求相关的值。
	//它可以帮助我们在多个 Goroutine 中有效地管理请求的生命周期，并提供了一种在处理函数中传递请求域相关数据和控制信号的标准方法。
	select {
	case <-ctx.Done():
		return true
	// 当上面没有收到消息 被阻塞的时候 那么就会执行default，表示还没被取消
	default:
		return false
	}
}

func cancel_2(cancelChan chan struct{}) {
	// 关闭同样是可以是阻塞状态channel 接收数据被唤醒执行下去，即自带的广播机制
	close(cancelChan)
}

// 空结构体 struct{} 在 Go 中被用作占位符，因为它不占用任何内存空间，但可以作为通道中的信号进行传递
// 通过发送一个取消指令（本次指只调用一次cancel_1(cancelChan)）那么也只会有一个协程收到消息然后释放
func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}

}
