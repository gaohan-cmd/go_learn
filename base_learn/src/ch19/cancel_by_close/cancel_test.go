package cancel_by_close

import (
	"fmt"
	"testing"
	"time"
)

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	// 同时启动5个协程
	for i := 0; i < 5; i++ {
		go func(i int, cancelChan chan struct{}) {
			for {
				if isCancel(cancelChan) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)
	}
	cancel_2(cancelChan)
	time.Sleep(time.Second * 1)

}

// 判断从收到的channel上是否收到了消息
func isCancel(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
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
// 由于struct{}不包含任何字段，因此它的实例在内存中不占据任何空间，所以使用它作为通道元素可以避免在通信时传输额外的数据，这对于单纯的信号传递而言是非常高效的。
// 使用空的struct{}作为通道类型是一种常见的方式来实现一个信号通道（signal channel）
func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}

}
